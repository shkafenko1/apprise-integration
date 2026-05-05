package container

import (
	"apprise-mvp/internal/config"
	"apprise-mvp/internal/email"
	appriseclient "apprise-mvp/pkg/apprise"
	"apprise-mvp/pkg/swagger"
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

type Container struct {
	emailHandler *email.EmailHandler
}

func NewContainer(cfg config.Config) *Container {
	slog.Debug("initializing container dependencies")
	appriseClient := appriseclient.New(cfg.AppriseBaseUrl)
	sEmail := email.NewEmailService(appriseClient, cfg.MailServerUrl)
	hEmail := email.NewEmailHandler(sEmail)

	return &Container{
		emailHandler: hEmail,
	}
}

func (c *Container) SetupRoutes(app *fiber.App) {
	email.RegisterRoutes(app, c.emailHandler)
	swagger.SwagInit(app)
}
