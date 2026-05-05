package appriseclient

import (
	"apprise-mvp/internal/storage/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	baseURL string
	key     string
	http    *http.Client
}

func New(baseURL, key string) *Client {
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		key:     key,
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Send(ctx context.Context, msg models.Email) error {
	url := fmt.Sprintf("%s/notify/%s", c.baseURL, c.key)

	payload, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal apprise message: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("failed to create apprise request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("network error sending to apprise: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("apprise returned %s", resp.Status)
	}

	return nil
}
