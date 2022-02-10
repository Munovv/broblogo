package config

import (
	"github.com/Munovv/broblogo/auth-service/config/middleware"
	"github.com/Munovv/broblogo/auth-service/config/mongo"
	"github.com/Munovv/broblogo/auth-service/config/server"
	"github.com/spf13/viper"
)

func Init(v *viper.Viper) error {
	viper.AddConfigPath("config")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}

func InitServer(v *viper.Viper) server.ServerConf {
	return server.ServerConf{
		Port:         server.GetServerPort(v),
		ReadTimeout:  server.GetServerReadTimeout(v),
		WriteTimeout: server.GetServerWriteTimeout(v),
	}
}

func InitMongoDb(v *viper.Viper) *mongo.MongoConfig {
	return &mongo.MongoConfig{
		Uri:        mongo.GetMongoDbUri(v),
		Name:       mongo.GetMongoDbName(v),
		Collection: mongo.GetCollectionName(v),
	}
}

func InitMiddleware(v *viper.Viper) *middleware.Middleware {
	return &middleware.Middleware{
		AccessToken: middleware.GetAccessToken(v),
		HashSalt:    middleware.GetHashSalt(v),
	}
}
