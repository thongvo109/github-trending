package repository

import (
	"context"
	"github-trending/model"
	"github-trending/model/req"
)

type UserRepo interface {
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	SaveUser(context context.Context, user model.User) (model.User, error)
	SelectUserById(context context.Context, userId string) (model.User, error)
	UpdateUser(ctx context.Context,user model.User)(model.User,error)
}
