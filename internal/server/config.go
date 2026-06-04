package http_server

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Address         string        `envconfig:"ADDRESS" default:":8083"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := envconfig.Process("SERVER", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewConfigMust() *Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}
