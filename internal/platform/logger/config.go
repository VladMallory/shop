package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Folder string `envconfig:"LOGGER_FOLDER" required:"true"`
	// DEBUG (-4), INFO (0), WARN (4) и ERROR (8).
	LogLevel int `envconfig:"LOGGER_LOG_LEVEL" default:"-4"`
}

func NewConfig() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to process env logger: %w", err)
	}

	return cfg, nil
}

// func NewConfigMust() Config {
// 	cfg, err := NewConfig()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	return cfg
// }

// .env
// LOGGER_FOLDER
// LOGGER_LOG_LEVEL
