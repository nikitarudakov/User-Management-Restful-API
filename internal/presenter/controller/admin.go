package controller

import (
	"crypto/subtle"
	"git.foxminded.ua/foxstudent106092/user-management/config"
	"git.foxminded.ua/foxstudent106092/user-management/internal/business/usecase/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type AdminController struct {
	ur       repository.UserRepository
	pr       repository.ProfileRepository
	cfgAdmin *config.Admin
	UserEndpointsHandler
	AuthEndpointHandler
}

func NewAdminController(ur repository.UserRepository, pr repository.ProfileRepository,
	cfgAdmin *config.Admin, uh UserEndpointsHandler, ah AuthEndpointHandler) *AdminController {

	return &AdminController{ur, pr, cfgAdmin,
		uh, ah}
}

func (ac *AdminController) InitRoutes(e *echo.Echo) {
	admin := e.Group("/admin")

	roles := []string{"admin", "moderator"}

	ac.InitAuthMiddleware(admin, roles)

	admin.GET("/users/profiles", func(ctx echo.Context) error {
		return ac.GetUserProfiles(ctx)
	})

	admin.PUT("/users/profiles/:username/modify", func(ctx echo.Context) error {
		return ac.ModifyUserProfile(ctx)
	})

	admin.DELETE("/users/profiles/:username/delete", func(ctx echo.Context) error {
		return ac.DeleteUserProfile(ctx)
	})
}

func (ac *AdminController) Auth(username, password string, adminCfg *config.Admin) (bool, error) {
	if subtle.ConstantTimeCompare([]byte(username), []byte(adminCfg.Username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(adminCfg.Password)) == 1 {

		return true, nil
	}

	return false, nil
}

func (ac *AdminController) GetUserProfiles(ctx echo.Context) error {
	var page int64 = 1

	pageStr := ctx.QueryParam("page")
	if pageStr != "" {
		parsedPage, err := strconv.ParseInt(ctx.QueryParam("page"), 10, 64)
		if err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}

		page = parsedPage
	}

	result, err := ac.pr.ListUserProfiles(page)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, result)
}

func (ac *AdminController) ModifyUserProfile(ctx echo.Context) error {
	err := ac.UpdateUserProfile(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (ac *AdminController) DeleteUserProfile(ctx echo.Context) error {
	authUsername := ctx.Param("username")

	err := ac.pr.Delete(authUsername)
	if err != nil {
		return err
	}

	err = ac.ur.Delete(authUsername)
	if err != nil {
		return err
	}

	return nil
}