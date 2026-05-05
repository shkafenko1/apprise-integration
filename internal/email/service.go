package email

import (
	"apprise-mvp/internal/storage/dtos"
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

func (es *EmailService) SendEmail(ctx context.Context, to, subject, body string) error {
	if to == "" || subject == "" || body == "" {
		return errors.New("empty fields")
	}

	fullUrl := fmt.Sprintf("%s&to=%s", es.smtpUrl, to)

	message := dtos.EmailToSend{
		Title: subject,
		Body:  body,
		URLs:  fullUrl,
	}

	return es.client.Send(ctx, message)
}
