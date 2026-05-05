package email

import (
	"apprise-mvp/internal/storage/models"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

type EmailHandler struct {
	service *EmailService
}

func NewEmailHandler(service *EmailService) *EmailHandler {
	return &EmailHandler{
		service: service,
	}
}

func (h *EmailHandler) SendEmail(ctx fiber.Ctx) error {
	var req models.Email

	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid json")
	}

	err := h.service.SendEmail(ctx.Context(), req.Subject, req.Body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	// On success
	return ctx.SendStatus(fiber.StatusNoContent)
}
