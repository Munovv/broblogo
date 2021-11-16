package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/Munovv/broblogo/AuthService/pkg/user"
	"github.com/Munovv/broblogo/AuthService/pkg/user/model"
	"github.com/Munovv/broblogo/AuthService/pkg/user/repository/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type Service struct {
	repository *mongo.Repository
	hashSalt   string
}

func NewUserService(r *mongo.Repository, hashSalt string) *Service {
	return &Service{
		repository: r,
		hashSalt:   hashSalt,
	}
}

func (s *Service) GetUser(ctx context.Context, guid string) (*model.User, error) {
	u, err := s.repository.GetUser(ctx, guid)
	if err != nil {
		return nil, user.ErrUserNotFound
	}

	return u, nil
}

func (s *Service) CreateUser(ctx context.Context, username, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(s.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	guid, err := generateGuid()
	if err != nil {
		return "", err
	}

	u := &model.User{
		Username: username,
		Password: password,
		GUID:     guid,
	}
	err = s.repository.CreateUser(ctx, u)
	if err != nil {
		return "", err
	}

	return guid, nil
}

func generateGuid() (string, error) {
	b, err := uuid.New()
	if err != nil {
		return "", err
	}

	guid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return guid, nil
}
