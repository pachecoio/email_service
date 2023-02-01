package mailgun

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/pachecoio/email_service/domain"
	"time"
)

type Client struct {
	mg mailgun.Mailgun
}

func NewClient(domain string, apiKey string) *Client {
	return &Client{
		mg: mailgun.NewMailgun(domain, apiKey),
	}
}

func (c *Client) Send(payload *domain.EmailPayload) error {
	message := c.mg.NewMessage(
		payload.From,
		payload.Subject,
		payload.Body,
		payload.To,
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := c.mg.Send(ctx, message)
	fmt.Printf("ID: %s Resp: %s\n", id, resp)

	if err != nil {
		return domain.EmailSendError{
			Message: err.Error(),
		}
	}
	return nil
}
