package services

import (
	"context"
	"errors"
	"go-curd-demo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u UserServiceImpl) CreateUser(user *models.User) (_ error) {
	_, error := u.usercollection.InsertOne(u.ctx, user)
	return error
}

func (u UserServiceImpl) GetUser(name string) (_ *models.User, _ error) {
	var user *models.User
	query := bson.D{bson.E{Key: "name", Value: name}}
	error := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, error
}

func (u UserServiceImpl) GetAll() (_ []*models.User, _ error) {
	var userslice []*models.User
	query := bson.M{}
	userCursor, error := u.usercollection.Find(u.ctx, query)

	if error != nil {
		return userslice, error
	}

	for userCursor.Next(u.ctx) {
		var user *models.User
		if err := userCursor.Decode(&user); err != nil {
			return userslice, err
		}
		userslice = append(userslice, user)
	}

	return userslice, error
}

func (u UserServiceImpl) UpateUser(user *models.User) (_ error) {
	query := bson.D{bson.E{Key: "userName", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "userName", Value: user.Name}, bson.E{Key: "userAddress", Value: user.Address}, bson.E{Key: "userAge", Value: user.Age}}}}
	result, err := u.usercollection.UpdateOne(u.ctx, query, update)
	if result.MatchedCount != 1 {
		//return errors.New("no match document found")
		return err
	}
	return nil
}

func (u UserServiceImpl) DeleteUser(name string) (_ error) {
	query := bson.D{bson.E{Key: "userName", Value: name}}
	result, err := u.usercollection.DeleteOne(u.ctx, query)
	if err != nil {
		return err
	}
	if result.DeletedCount != 1 {
		return errors.New("no match Found to delete")
	}
	return nil
}
