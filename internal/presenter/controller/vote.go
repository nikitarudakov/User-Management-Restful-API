package controller

import (
	"errors"
	"fmt"
	"git.foxminded.ua/foxstudent106092/user-management/internal/business/model"
	"git.foxminded.ua/foxstudent106092/user-management/internal/infrastructure/appErrors/repoerr"
	"git.foxminded.ua/foxstudent106092/user-management/internal/presenter/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type VoteController struct {
	voteUsecase VoteManager
	AuthEndpointHandler
}

func NewVoteController(vc VoteManager, ac AuthEndpointHandler) *VoteController {
	return &VoteController{vc, ac}
}

type VoteManager interface {
	StoreVote(v *model.Vote) (*repository.VoteInsertResult, error)
	RetractVote(u *model.Update, sender string) error
	GetRating(target string) (*model.Rating, error)
}

func (vc *VoteController) InitRoutes(g *echo.Group) {
	roles := []string{"admin", "moderator", "user"}

	vc.InitAuthMiddleware(g, roles)

	g.GET("/profiles/:target/rating", func(ctx echo.Context) error {
		return vc.GetRating(ctx)
	})

	g.PUT("/profiles/:target/vote", func(ctx echo.Context) error {
		return vc.Vote(ctx)
	})

	g.DELETE("/profiles/:target/vote/retract", func(ctx echo.Context) error {
		return vc.RetractVote(ctx)
	})
}

func (vc *VoteController) Vote(ctx echo.Context) error {
	vote, err := vc.parseUserVote(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := vc.voteUsecase.StoreVote(vote)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, result)
}

func (vc *VoteController) RetractVote(ctx echo.Context) error {
	sender := fmt.Sprintf("%v", ctx.Get("username"))
	target := ctx.Param("target")

	var update = &model.Update{Nickname: target}

	err := vc.voteUsecase.RetractVote(update, sender)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (vc *VoteController) GetRating(ctx echo.Context) error {
	target := ctx.Param("target")

	rating, err := vc.voteUsecase.GetRating(target)
	if err != nil {
		var calcRatingUserError *repoerr.CalcRatingUserError
		if errors.As(err, &calcRatingUserError) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, rating)
}

func (vc *VoteController) parseUserVote(ctx echo.Context) (*model.Vote, error) {
	sender := fmt.Sprintf("%v", ctx.Get("username"))
	target := ctx.Param("target")

	if sender == target {
		return nil, errors.New("self-voting is forbidden (bad request)")
	}

	voteStr := ctx.FormValue("vote")
	vote, err := strconv.Atoi(voteStr)
	if err != nil || (vote != -1 && vote != 1) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	now := time.Now().Unix()

	voteObj := model.Vote{
		Sender:  sender,
		Target:  target,
		Vote:    vote,
		VotedAt: now,
	}

	return &voteObj, nil
}