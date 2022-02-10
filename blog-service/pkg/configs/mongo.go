package configs

import "github.com/spf13/viper"

type Mongo struct {
	Uri        string
	Name       string
	Collection string
}

func (ctg *Mongo) getConfig(v *viper.Viper) *Mongo {
	return &Mongo{
		Uri:        getMongoDbUri(v),
		Name:       getMongoDbName(v),
		Collection: getCollectionName(v),
	}
}

func getMongoDbUri(v *viper.Viper) string {
	return v.GetString("mongo.uri")
}

func getMongoDbName(v *viper.Viper) string {
	return v.GetString("mongo.name")
}

func getCollectionName(v *viper.Viper) string {
	return v.GetString("mongo.collection")
}
