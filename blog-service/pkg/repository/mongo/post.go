package mongo

import (
	"context"
	"github.com/Munovv/broblogo/blog-service/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID  primitive.ObjectID `bson:"userId"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
	GUID    string             `bson:"guid"`
}

type PostRepository struct {
	db *mongo.Collection
}

func NewPostRepository(db *mongo.Database, collection string) *PostRepository {
	return &PostRepository{
		db: db.Collection(collection),
	}
}

func (r *PostRepository) CreatePost(ctx context.Context, p *models.Post) error {
	mongoModel := toMongoPost(p)

	res, err := r.db.InsertOne(ctx, mongoModel)
	if err != nil {
		return err
	}

	p.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (r *PostRepository) GetPosts(ctx context.Context, uid string) ([]*models.Post, error) {
	hexUid, _ := primitive.ObjectIDFromHex(uid)
	res, err := r.db.Find(ctx, bson.M{
		"userId": hexUid,
	})
	defer res.Close(ctx)

	if err != nil {
		return nil, err
	}

	out := make([]*Post, 0)

	for res.Next(ctx) {
		post := new(Post)
		err := res.Decode(post)
		if err != nil {
			return nil, err
		}

		out = append(out, post)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return toPostCollection(out), nil
}

func (r *PostRepository) GetPost(ctx context.Context, guid string, uid string) (*models.Post, error) {
	post := new(Post)
	err := r.db.FindOne(ctx, bson.M{
		"guid":   guid,
		"userId": uid,
	}).Decode(post)

	if err != nil {
		return nil, err
	}

	return toModelPost(post), nil
}

func (r *PostRepository) EditPost(ctx context.Context, p *models.Post) error {
	filter := bson.M{
		"guid":   p.GUID,
		"userId": p.UserId,
	}
	_, err := r.db.UpdateOne(ctx, filter, p)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) RemovePost(ctx context.Context, guid string, uid string) error {
	_, err := r.db.DeleteOne(ctx, bson.M{
		"userId": uid,
		"guid":   guid,
	})

	if err != nil {
		return err
	}

	return nil
}

func toMongoPost(p *models.Post) *Post {
	uid, _ := primitive.ObjectIDFromHex(p.UserId)
	return &Post{
		UserID:  uid,
		Title:   p.Title,
		Content: p.Content,
		GUID:    p.GUID,
	}
}

func toModelPost(p *Post) *models.Post {
	return &models.Post{
		ID:      p.ID.Hex(),
		UserId:  p.UserID.Hex(),
		Title:   p.Title,
		Content: p.Content,
		GUID:    p.GUID,
	}
}

func toPostCollection(ps []*Post) []*models.Post {
	out := make([]*models.Post, len(ps))

	for key, p := range ps {
		out[key] = toModelPost(p)
	}

	return out
}
