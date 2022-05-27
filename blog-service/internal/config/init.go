package config

import "github.com/spf13/viper"

type Config struct {
	Server *Server
	Mongo  *Mongo
}

func initReader() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
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
