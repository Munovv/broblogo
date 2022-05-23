package service

import (
	"context"
	"github.com/Munovv/broblogo/partner-service/internal/models"
)

type Repository interface {
	CreatePartner(ctx context.Context, partner *models.Partner) error
	GetPartner(ctx context.Context, id string) (*models.Partner, error)
	GetPartners(ctx context.Context) ([]*models.Partner, error)
	DeletePartner(ctx context.Context, id string) error
}

type service struct {
	repository Repository
}

func (s *service) CreatePartner(ctx context.Context, input *models.CreatePartnerInput) error {
	partner := &models.Partner{
		Name:        input.Name,
		Description: input.Description,
		Location:    input.Location,
	}

	return s.repository.CreatePartner(ctx, partner)
}

func (s *service) GetPartner(ctx context.Context, id string) (*models.Partner, error) {
	return s.repository.GetPartner(ctx, id)
}

func (s *service) GetPartners(ctx context.Context) ([]*models.Partner, error) {
	return s.repository.GetPartners(ctx)
}

func (s *service) DeletePartner(ctx context.Context, id string) error {
	return s.repository.DeletePartner(ctx, id)
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}
