package repository

import (
	"context"
	"github.com/Munovv/broblogo/internal/blog-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database interface {
	GetAbstractDatabase() *mongo.Database
}

type repository struct {
	db *mongo.Collection
}

func (r *repository) Create(ctx context.Context, post *model.Post) error {
	p := r.toMongoModel(post)
	res, err := r.db.InsertOne(ctx, p)
	if err != nil {
		return err
	}

	post.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (r *repository) Get(ctx context.Context, id string) (*model.Post, error) {
	p := new(model.PostMongo)
	err := r.db.FindOne(ctx, bson.M{
		"ID": id,
	}).Decode(p)

	if err != nil {
		return nil, err
	}

	return r.toModel(p), nil
}

func (r *repository) GetBy(ctx context.Context, filter interface{}) ([]model.Post, error) {
	var postsMongo []model.PostMongo
	var postModels []model.Post

	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		return postModels, err
	}

	if err = cursor.All(ctx, &postsMongo); err != nil {
		return postModels, err
	}

	for _, mongoObj := range postsMongo {
		postModels = append(postModels, *r.toModel(&mongoObj))
	}

	return postModels, nil
}

func (r *repository) Update(ctx context.Context, post *model.Post) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"ID": post.ID}, bson.M{
		"Title":       post.Title,
		"Description": post.Description,
		"Content":     post.Content,
	})

	return err
}

func (r *repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.DeleteOne(ctx, bson.M{
		"ID": id,
	})

	return err
}

func (r *repository) toMongoModel(post *model.Post) *model.PostMongo {
	return &model.PostMongo{
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		CreatedAt:   post.CreatedAt,
		UserId:      post.UserId,
	}
}

func (r *repository) toModel(post *model.PostMongo) *model.Post {
	return &model.Post{
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		CreatedAt:   post.CreatedAt,
		UserId:      post.UserId,
	}
}

func NewRepository(db Database, collection string) *repository {
	return &repository{
		db: (db.GetAbstractDatabase()).Collection(collection),
	}
}
