package services

import (
	"context"
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

func (u UserServiceImpl) UpateUser(_ *models.User) (_ error) {
	return nil
}

func (u UserServiceImpl) DeleteUser(_ *models.User) (_ error) {
	return nil
}
