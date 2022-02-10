package service

import (
	"context"
	"fmt"
	"github.com/Munovv/broblogo/blog-service/pkg/messages"
	"github.com/Munovv/broblogo/blog-service/pkg/models"
	"github.com/Munovv/broblogo/blog-service/pkg/repository"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) CreatePost(ctx context.Context, p *models.Post) error {
	guid, err := generateGuid()
	if err != nil {
		return nil
	}

	p.GUID = guid

	return s.repo.CreatePost(ctx, p)
}

func (s *PostService) GetPosts(ctx context.Context, uid string) ([]*models.Post, error) {
	ps, err := s.repo.GetPosts(ctx, uid)
	if err != nil {
		return nil, messages.ErrPostNotFound
	}

	return ps, nil
}

func (s *PostService) GetPost(ctx context.Context, guid string, uid string) (*models.Post, error) {
	p, err := s.repo.GetPost(ctx, guid, uid)
	if err != nil {
		return nil, messages.ErrPostNotFound
	}

	return p, nil
}

func (s *PostService) EditPost(ctx context.Context, p *models.Post) error {
	return s.repo.EditPost(ctx, p)
}

func (s *PostService) RemovePost(ctx context.Context, guid string, uid string) error {
	return s.repo.RemovePost(ctx, guid, uid)
}

func generateGuid() (string, error) {
	b, err := uuid.New()
	if err != nil {
		return "", err
	}

	guid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return guid, nil
}
