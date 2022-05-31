package config

import "time"

type Server struct {
	Port              string        `yaml:"port"`
	ReadTimeout       time.Duration `yaml:"read_timeout"`
	WriteTimeout      time.Duration `yaml:"write_timeout"`
	MaxWriteMegabytes int           `yaml:"max_write_megabytes"`
}
