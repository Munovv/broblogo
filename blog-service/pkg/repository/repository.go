package repository

import (
	"context"
	"github.com/Munovv/broblogo/blog-service/pkg/configs"
	"github.com/Munovv/broblogo/blog-service/pkg/models"
	mongoRepository "github.com/Munovv/broblogo/blog-service/pkg/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Post interface {
	CreatePost(ctx context.Context, post *models.Post) error
	GetPosts(ctx context.Context, userId string) ([]*models.Post, error)
	GetPost(ctx context.Context, guid string, userId string) (*models.Post, error)
	EditPost(ctx context.Context, post *models.Post) error
	RemovePost(ctx context.Context, guid string, userId string) error
}

type Repository struct {
	Post
}

func NewMongoRepository(db *mongo.Database, cfg *configs.Mongo) *Repository {
	return &Repository{
		Post: mongoRepository.NewPostRepository(db, cfg.Collection),
	}
}
