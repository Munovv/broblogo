package repository

import (
	"context"
	"github.com/Munovv/broblogo/auth-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Database, col string) *Repository {
	return &Repository{
		db: db.Collection(col),
	}
}

func (r Repository) CreateUser(ctx context.Context, u *model.User) error {
	uModel := toMongoUser(u)
	res, err := r.db.InsertOne(ctx, uModel)
	if err != nil {
		return err
	}

	u.Id = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (r Repository) GetUser(ctx context.Context, username, password string) (*model.User, error) {
	u := new(model.UserMongo)
	err := r.db.FindOne(ctx, bson.M{
		"username": username,
		"password": password,
	}).Decode(u)

	if err != nil {
		return nil, err
	}

	return toModelUser(u), nil
}

func toMongoUser(u *model.User) *model.UserMongo {
	return &model.UserMongo{
		Username: u.Username,
		Password: u.Password,
	}
}

func toModelUser(u *model.UserMongo) *model.User {
	return &model.User{
		Username: u.Username,
		Password: u.Password,
	}
}
