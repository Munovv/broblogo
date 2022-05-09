package config

import (
	"github.com/Munovv/broblogo/auth-service/internal/config/mongo"
	"github.com/Munovv/broblogo/auth-service/internal/config/server"
	"github.com/spf13/viper"
)

type Config struct {
	Server *server.ServerConf
	Mongo  *mongo.MongoConfig
}

func initReader() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func initServer(v *viper.Viper) *server.ServerConf {
	return &server.ServerConf{
		Port:         server.GetServerPort(v),
		ReadTimeout:  server.GetServerReadTimeout(v),
		WriteTimeout: server.GetServerWriteTimeout(v),
	}
}

func initMongoDb(v *viper.Viper) *mongo.MongoConfig {
	return &mongo.MongoConfig{
		Uri:        mongo.GetMongoDbUri(v),
		Name:       mongo.GetMongoDbName(v),
		Collection: mongo.GetCollectionName(v),
	}
}

func NewConfig() (*Config, error) {
	if err := initReader(); err != nil {
		return nil, err
	}

	return &Config{
		Server: initServer(viper.GetViper()),
		Mongo:  initMongoDb(viper.GetViper()),
	}, nil
}
