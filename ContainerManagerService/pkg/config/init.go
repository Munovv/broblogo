package config

import "github.com/spf13/viper"

type Config struct {
	Mongo      *Mongo
	Server     *Server
	Middleware *Middleware
}

func Init(v *viper.Viper) (*Config, error) {
	if err := readFile(v); err != nil {
		return &Config{}, err
	}

	cfgMongo := new(Mongo)
	cfgServer := new(Server)
	cfgMw := new(Middleware)

	return &Config{
		Mongo:      cfgMongo.getConfig(v),
		Server:     cfgServer.getConfig(v),
		Middleware: cfgMw.getConfig(v),
	}, nil
}

func readFile(v *viper.Viper) error {
	v.AddConfigPath("configs")
	v.SetConfigName("main")

	return v.ReadInConfig()
}
