package database_postgres

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	User     string        `envconfig:"USER" required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	Host     string        `envconfig:"HOST" required:"true"`
	Port     string        `envconfig:"PORT" required:"true"`
	Database string        `envconfig:"DATABASE" required:"true"`
	Timeout  time.Duration `envconfig:"TIMEOUT" required:"true"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("POSTGRES", cfg)
	if err != nil {
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
