package repository

import (
	"context"
	"errors"
	"fmt"
	"git.foxminded.ua/foxstudent106092/user-management/config"
	"git.foxminded.ua/foxstudent106092/user-management/internal/domain/model"
	"git.foxminded.ua/foxstudent106092/user-management/internal/usecase/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type profileRepository struct {
	db *mongo.Client
}

func NewProfileRepository(db *mongo.Client) repository.ProfileRepository {
	return &profileRepository{db: db}
}

func (pr *profileRepository) Create(p *model.Profile) (interface{}, error) {
	_, err := pr.Find(p)
	if err == nil {
		return nil, errors.New("profile with such nickname already exists")
	}

	now := time.Now().Unix()
	p.CreatedAt = &now
	p.UpdatedAt = &now

	result, err := pr.InsertMethod(p)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (pr *profileRepository) Update(p *model.Profile) error {
	now := time.Now().Unix()
	p.UpdatedAt = &now

	err := pr.UpdateMethod(p)
	if err != nil {
		return err
	}

	return nil
}

func (pr *profileRepository) Find(p *model.Profile) (*model.Profile, error) {
	if err := pr.FindMethod(p).Decode(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (pr *profileRepository) FindMethod(p *model.Profile) *mongo.SingleResult {
	coll := pr.getCollection(p.TableName())

	filter := bson.M{"nickname": p.Nickname}

	result := coll.FindOne(context.TODO(), filter)

	return result
}

func (pr *profileRepository) InsertMethod(p *model.Profile) (interface{}, error) {
	coll := pr.getCollection(p.TableName())

	result, err := coll.InsertOne(context.TODO(), p)
	if err != nil {
		return nil, fmt.Errorf("error updating/inserting user data: %w", err)
	}

	return result.InsertedID, nil
}

func (pr *profileRepository) UpdateMethod(p *model.Profile) error {
	coll := pr.getCollection(p.TableName())

	filter := bson.M{"nickname": p.Nickname}
	update := bson.M{"$set": bson.M{
		"first_name": p.FirstName,
		"last_name":  p.LastName,
		"updated_at": time.Now().Unix(),
	}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("error updating/inserting user data: %w", err)
	}

	if result.MatchedCount == 0 {
		return errors.New("there is no profile to update")
	}

	return nil
}

func (pr *profileRepository) getCollection(colName string) *mongo.Collection {
	return pr.db.Database(config.C.Database.Name).Collection(colName)
}
