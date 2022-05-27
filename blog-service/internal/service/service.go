package service

import (
	"context"
	"github.com/Munovv/broblogo/blog-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Repository interface {
	Create(ctx context.Context, post *model.Post) error
	Get(ctx context.Context, id string) (*model.Post, error)
	GetBy(ctx context.Context, filter interface{}) ([]model.Post, error)
	Update(ctx context.Context, post *model.Post) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	repo Repository
}

func (s *service) CreateItem(ctx context.Context, input model.CreatePostInput, userId string) error {
	post := &model.Post{
		UserId:      userId,
		Title:       input.Title,
		Description: input.Description,
		Content:     input.Content,
		CreatedAt:   time.Now(),
	}

	return s.repo.Create(ctx, post)
}

func (s *service) GetItem(ctx context.Context, id string) (*model.Post, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) GetItems(ctx context.Context, userId string) ([]model.Post, error) {
	return s.repo.GetBy(ctx, bson.M{
		"UserId": userId,
	})
}

func (s *service) UpdateItem(ctx context.Context, id string, input model.UpdatePostInput) error {
	post := &model.Post{
		ID:          id,
		Title:       input.Title,
		Description: input.Title,
		Content:     input.Content,
	}

	return s.repo.Update(ctx, post)
}

func (s *service) DeleteItem(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}
