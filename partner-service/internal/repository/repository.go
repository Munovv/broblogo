package repository

import (
	"context"
	"github.com/Munovv/broblogo/partner-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Collection
}

func (r *repository) CreatePartner(ctx context.Context, partner *models.Partner) error {
	resp, err := r.db.InsertOne(ctx, partner)
	if err != nil {
		return err
	}

	partner.Id = resp.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (r *repository) GetPartner(ctx context.Context, id string) (*models.Partner, error) {
	var partner models.PartnerMongo

	idFromHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	resp := r.db.FindOne(ctx, bson.M{"_id": idFromHex}).Decode(&partner)

	return r.toModel(&partner), resp
}

func (r *repository) GetPartners(ctx context.Context) ([]*models.Partner, error) {
	partners := make([]*models.PartnerMongo, 0)

	cursor, err := r.db.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		partner := new(models.PartnerMongo)
		if err := cursor.Decode(partner); err != nil {
			return nil, err
		}

		partners = append(partners, partner)
	}

	return r.toModels(partners), nil
}

func (r *repository) DeletePartner(ctx context.Context, id string) error {
	idFromHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.db.DeleteOne(ctx, bson.M{"_id": idFromHex})

	return err
}

func (r *repository) toMongoModel(partner *models.Partner) *models.PartnerMongo {
	return &models.PartnerMongo{
		Name:        partner.Name,
		Location:    partner.Location,
		Description: partner.Description,
	}
}

func (r *repository) toModel(partner *models.PartnerMongo) *models.Partner {
	return &models.Partner{
		Id:          partner.Id.Hex(),
		Name:        partner.Name,
		Location:    partner.Location,
		Description: partner.Description,
	}
}

func (r *repository) toModels(partners []*models.PartnerMongo) []*models.Partner {
	partnerModels := make([]*models.Partner, 0)

	for _, p := range partners {
		partnerModels = append(partnerModels, r.toModel(p))
	}

	return partnerModels
}

func NewRepository(db *mongo.Database, collection string) *repository {
	return &repository{
		db: db.Collection(collection),
	}
}
