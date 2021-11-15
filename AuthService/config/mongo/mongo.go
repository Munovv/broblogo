package mongo

import "github.com/spf13/viper"

type MongoConfig struct {
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
