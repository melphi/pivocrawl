package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ApiPort int `envconfig:"API_PORT" default:"8080"`
}

func New() *Config {
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("Error loading environment properties: [%s].", err)
	}
	return &cfg
}
