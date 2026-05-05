package email

import "github.com/gofiber/fiber/v3"

type EmailRouter struct {
	handler *EmailHandler
}

func NewEmailRouter(handler *EmailHandler) *EmailRouter {
	return &EmailRouter{
		handler: handler,
	}
}

func RegisterRoutes(router fiber.Router, h *EmailHandler) {
	router.Post("/send", h.SendEmail)
}
