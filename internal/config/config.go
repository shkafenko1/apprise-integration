package config

import (
	"apprise-mvp/pkg/envparse"
	"log/slog"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort       string
	AppriseBaseUrl string
	MailServerUrl  string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
	}

	return Config{
		HTTPPort:       envparse.GetEnv("APP_PORT", "8080"),
		AppriseBaseUrl: envparse.GetEnv("APPRISE_API_URL", "http://localhost:8000"),
		MailServerUrl:  envparse.GetEnv("MAIL_SERVER_URL", ""),
	}
}
