package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	yamlFormat = "yml"
)

type Config struct {
	Mongo  Mongo  `yaml:"mongodb" validate:"required"`
	Server Server `yaml:"server" validate:"required"`
}

func buildPath(path, name string) string {
	return fmt.Sprintf("%s/%s.%s", path, name, yamlFormat)
}

func NewConfig(path, name string) (*Config, error) {
	var cfg Config

	fullPath := buildPath(path, name)
	err := cleanenv.ReadConfig(fullPath, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
