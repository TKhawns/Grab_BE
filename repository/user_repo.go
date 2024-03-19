package repository

import (
	"context"
	"go_grab/model"
	"go_grab/model/request"
)

type UserRepo interface {
	CheckSignIn(context context.Context, signInReq request.RequestSignIn) (model.User, error)
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckIfExist(context context.Context, signUpReq request.RequestSignUp) error
}
