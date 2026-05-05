package app

import (
	"apprise-mvp/internal/config"
	"apprise-mvp/internal/container"
	"apprise-mvp/pkg/logger"
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

func App() {
	logger.InitLogger()
	slog.Info("starting application")

	cfg := config.Load()
	slog.Info("configuration loaded")

	c := container.NewContainer(cfg)
	slog.Info("container initialized")

	app := fiber.New()
	c.SetupRoutes(app)
	slog.Info("routes setup completed")

	slog.Info("starting server", "port", cfg.HTTPPort)
	if err := app.Listen(":" + cfg.HTTPPort); err != nil {
		slog.Error("server failed to start", "error", err)
	}
}
