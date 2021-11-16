package mongo

import (
	"context"
	"github.com/Munovv/broblogo/AuthService/pkg/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	GUID     string             `bson:"guid"`
}

type Repository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collection string) *Repository {
	return &Repository{
		db: db.Collection(collection),
	}
}

func (r Repository) CreateUser(ctx context.Context, u *model.User) error {
	uModel := toMongoUser(u)
	res, err := r.db.InsertOne(ctx, uModel)
	if err != nil {
		return err
	}

	u.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (r Repository) GetUser(ctx context.Context, guid string) (*model.User, error) {
	u := new(User)
	err := r.db.FindOne(ctx, bson.M{
		"guid": guid,
	}).Decode(u)

	if err != nil {
		return nil, err
	}

	return toModelUser(u), nil
}

func toMongoUser(u *model.User) *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
		GUID:     u.GUID,
	}
}

func toModelUser(u *User) *model.User {
	return &model.User{
		Username: u.Username,
		Password: u.Password,
		GUID:     u.GUID,
	}
}
