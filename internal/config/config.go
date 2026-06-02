package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	JWTToken string
}

// Config возвращает конфигурацию env.
func Config() *Configs {
	if err := godotenv.Load(); err != nil {
		log.Fatal("не грузится конфиг")
	}

	return &Configs{
		JWTToken: os.Getenv("jwt_token"),
	}
}
