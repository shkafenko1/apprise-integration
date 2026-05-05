package config

import (
	"apprise-mvp/pkg/envparse"
	"log/slog"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort       string
	AppriseBaseUrl string
	AppriseKey     string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
	}

	return Config{
		HTTPPort:       envparse.GetEnv("APP_PORT", "8080"),
		AppriseBaseUrl: envparse.GetEnv("APPRISE_BASE_URL", "http://localhost:8000"),
		AppriseKey:     envparse.GetEnv("APPRISE_KEY", "required"),
	}
}
