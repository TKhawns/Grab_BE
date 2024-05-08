package repository

import (
	"context"
	"go_grab/model"
	"go_grab/model/request"
)

type AdminRepo interface {
	GetUserModel(context context.Context) ([]model.Model, error)
	UpdateStatus(context context.Context, req request.RequestUpdateStatus) error
}
