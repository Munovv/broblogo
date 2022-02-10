package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Mongo      *Mongo
	Server     *Server
	Middleware *Middleware
}

func Init(v *viper.Viper) (*Config, error) {
	if err := readFile(v); err != nil {
		return &Config{}, err
	}

	return &Config{
		Mongo:      new(Mongo).getConfig(v),
		Server:     new(Server).getConfig(v),
		Middleware: new(Middleware).getConfig(v),
	}, nil
}

func readFile(v *viper.Viper) error {
	viper.AddConfigPath("./pkg/configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
