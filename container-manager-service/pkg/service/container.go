package service

import (
	"context"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/model"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/repository"
)

type ContainerService struct {
	repo repository.Container
}

func NewContainerService(repo repository.Container) *ContainerService {
	return &ContainerService{
		repo: repo,
	}
}

func (s *ContainerService) Create(ctx context.Context, cont *model.Container) (string, error) {
	return "", nil
}

func (s *ContainerService) Get(ctx context.Context, guid string) (*model.Container, error) {
	return nil, nil
}

func (s *ContainerService) Delete(ctx context.Context, guid string) error {
	return nil
}
