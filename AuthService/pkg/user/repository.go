package user

import (
	"context"
	"github.com/Munovv/broblogo/GoMainService/pkg/user/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, guid string) (*model.User, error)
}
