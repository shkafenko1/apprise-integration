package email

import (
	"apprise-mvp/internal/storage/models"
	"encoding/json"
	"log/slog"

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

	slog.InfoContext(ctx.Context(), "handling send email request")

	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		slog.ErrorContext(ctx.Context(), "failed to unmarshal send email request", "error", err, "status", fiber.StatusBadRequest)
		return fiber.NewError(fiber.StatusBadRequest, "invalid json")
	}

	err := h.service.SendEmail(ctx.Context(), req.Receiver, req.Subject, req.Body)
	if err != nil {
		slog.ErrorContext(ctx.Context(), "failed to send email via service", "error", err, "status", fiber.StatusBadGateway)
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	slog.InfoContext(ctx.Context(), "email sent successfully", "status", fiber.StatusNoContent)

	return ctx.SendStatus(fiber.StatusNoContent)
}
