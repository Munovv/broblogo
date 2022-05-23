package config

type Mongo struct {
	Connection string `yaml:"url" validate:"required"`
	Database   string `yaml:"database" validate:"required"`
	Collection string `yaml:"collection" validate:"required"`
}
