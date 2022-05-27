package config

import "github.com/spf13/viper"

type Mongo struct {
	Uri        string
	Name       string
	Collection string
}

func GetMongoDbUri(v *viper.Viper) string {
	return viper.GetString("mongo.uri")
}

func GetMongoDbName(v *viper.Viper) string {
	return viper.GetString("mongo.name")
}

func GetCollectionName(v *viper.Viper) string {
	return viper.GetString("mongo.collection")
}

func initMongoDb(v *viper.Viper) *Mongo {
	return &Mongo{
		Uri:        GetMongoDbUri(v),
		Name:       GetMongoDbName(v),
		Collection: GetCollectionName(v),
	}
}
