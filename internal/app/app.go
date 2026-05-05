package app

import (
	"apprise-mvp/internal/config"
	"apprise-mvp/internal/container"

	"github.com/gofiber/fiber/v3"
)

func App() {
	cfg := config.Load()
	c := container.NewContainer(cfg)
	app := fiber.New()
	c.SetupRoutes(app)

	app.Listen(":8080")
}
