package core_logger

import (
	"fmt"
	"log/slog"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Folder   string     `envconfig:"FOLDER" required:"true"`
	LogLevel slog.Level `envconfig:"LOG_LEVEL" default:"0"` //DEBUG (-4), INFO (0), WARN (4) и ERROR (8).
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("LOGGER", cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to process env logger: %w", err)
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
