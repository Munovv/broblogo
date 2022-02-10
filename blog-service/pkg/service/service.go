package service

import (
	"context"
	"github.com/Munovv/broblogo/blog-service/pkg/models"
	"github.com/Munovv/broblogo/blog-service/pkg/repository"
)

type Service struct {
	Post
}

type Post interface {
	CreatePost(ctx context.Context, post *models.Post) error
	GetPosts(ctx context.Context, userId string) ([]*models.Post, error)
	GetPost(ctx context.Context, guid string, userId string) (*models.Post, error)
	EditPost(ctx context.Context, post *models.Post) error
	RemovePost(ctx context.Context, guid string, userId string) error
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Post: NewPostService(repo.Post),
	}
}
