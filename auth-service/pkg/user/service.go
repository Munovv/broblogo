package user

import (
	"context"
	"github.com/Munovv/broblogo/auth-service/pkg/user/model"
	"github.com/Munovv/broblogo/auth-service/pkg/user/repository/mongo"
)

const CtxUserKey = "user"

type Service interface {
	NewUserService(r *mongo.Repository, hashSalt string) *Service
	GetUser(ctx context.Context, guid string) (*model.User, error)
	CreateUser(ctx context.Context, username, password string) error
}
