package email

import (
	"apprise-mvp/internal/storage/models"
	appriseclient "apprise-mvp/pkg/apprise"
	"context"
	"errors"
)

type EmailService struct {
	client *appriseclient.Client
}

func NewEmailService(client *appriseclient.Client) *EmailService {
	return &EmailService{
		client: client,
	}
}

func (es *EmailService) SendEmail(ctx context.Context, subject, body string) error {
	if subject == "" || body == "" {
		return errors.New("invalid email subject or body")
	}
	message := models.Email{
		Subject: subject,
		Body:    body,
	}

	return es.client.Send(ctx, message)
}
