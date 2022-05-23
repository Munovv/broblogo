package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Partner struct {
	Id          string
	Name        string
	Location    string
	Description string
}

type PartnerMongo struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Location    string             `bson:"location"`
	Description string             `bson:"description"`
}
