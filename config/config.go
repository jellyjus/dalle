package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port     int `default:"80"`
	Postgres Postgres
	OpenAI   OpenAI
}

type Postgres struct {
	Host     string `split_words:"true" default:"localhost"`
	Port     int    `split_words:"true" default:"5432"`
	User     string `required:"true" split_words:"true"`
	Password string `required:"true" split_words:"true"`
	Database string `required:"true" split_words:"true"`
}

type OpenAI struct {
	APIKey string `required:"true" split_words:"true"`
}

func GetConfig() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
