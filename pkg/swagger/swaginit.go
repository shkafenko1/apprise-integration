package swagger

import (
	"log/slog"

	"github.com/gofiber/contrib/v3/swaggerui"
	"github.com/gofiber/fiber/v3"
)

func SwagInit(app *fiber.App) {

	slog.Info("initializing swagger")

	app.Use(swaggerui.New(swaggerui.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Apprise MVP API Docs",
	}))
}
