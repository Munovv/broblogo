package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/Munovv/broblogo/auth-service/auth-service/errors"
	"github.com/Munovv/broblogo/auth-service/auth-service/model"
	"github.com/Munovv/broblogo/auth-service/auth-service/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "8dhadsb7nSDkWWq"
	signingKey = "7ashdsaaudads12casdj"
	tokenTTL   = 24 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

type Service struct {
	repo *repository.Repository
}

func (s *Service) CreateUser(ctx context.Context, username, password string) error {
	u := &model.User{
		Username: username,
		Password: s.generateHash(password),
	}

	return s.repo.CreateUser(ctx, u)
}

func (s *Service) GetUser(ctx context.Context, username, password string) (*model.User, error) {
	u, err := s.repo.GetUser(ctx, username, s.generateHash(password))
	if err != nil {
		return nil, errors.UserNotFound
	}

	return u, nil
}

func (s *Service) GenerateToken(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.GetUser(ctx, username, s.generateHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Service) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.InvalidSigningMethod
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.InvalidTokenType
	}

	return claims.UserId, nil
}

func (s *Service) generateHash(password string) string {
	pwd := sha256.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(salt))

	return fmt.Sprintf("%x", pwd.Sum(nil))
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
