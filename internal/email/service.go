package email

import (
	"apprise-mvp/internal/storage/dtos"
	"apprise-mvp/internal/storage/models"
	appriseclient "apprise-mvp/pkg/apprise"
	"context"
	"errors"
	"fmt"
)

type EmailService struct {
	client  *appriseclient.Client
	smtpUrl string
}

func NewEmailService(client *appriseclient.Client, smtpUrl string) *EmailService {
	return &EmailService{
		client:  client,
		smtpUrl: smtpUrl,
	}
}

func (es *EmailService) SendEmail(ctx context.Context, received models.Email) error {
	if received.Receiver == "" || received.Subject == "" || received.Body == "" {
		return errors.New("empty fields")
	}

	fullUrl := fmt.Sprintf("%s&to=%s", es.smtpUrl, received.Receiver)

	message := dtos.EmailToSend{
		Title:  received.Subject,
		Body:   received.Body,
		URLs:   fullUrl,
		Format: "html",
	}

	return es.client.Send(ctx, message)
}
