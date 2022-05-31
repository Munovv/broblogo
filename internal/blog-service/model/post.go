package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	ID          string
	UserId      string
	Title       string
	Description string
	Content     string
	CreatedAt   time.Time
}

type PostMongo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserId      string
	Title       string
	Description string
	Content     string
	CreatedAt   time.Time
}
