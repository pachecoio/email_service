package services

import (
	"github.com/pachecoio/email_service/adapters"
	"github.com/pachecoio/email_service/domain"
	"testing"
)

func TestSendEmail(t *testing.T) {
	type args struct {
		client  adapters.AbstractEmailClient
		payload *domain.EmailPayload
	}
	tests := []struct {
		name string
		args args
		res  domain.EmailSentEvent
		err  error
	}{
		// Test cases
		{
			name: "Test SendEmail with valid payload and client",
			args: args{
				client: &EmailClient{},
				payload: &domain.EmailPayload{
					From:    "hi@pacheco.io",
					To:      "thiagodelimapacheco@gmail.com",
					Subject: "Hello, World!",
					Body:    "<p>This is a test email</p>",
				},
			},
			res: domain.EmailSentEvent{
				From:    "hi@pacheco.io",
				To:      "thiagodelimapacheco@gmail.com",
				Subject: "Hello, World!",
				Body:    "<p>This is a test email</p>",
			},
			err: nil,
		},
		{
			name: "Test SendEmail with invalid payload and client",
			args: args{
				client: &EmailClientThatFails{},
				payload: &domain.EmailPayload{
					From:    "hi@pacheco.io",
					To:      "thiagodelimapacheco@gmail.com",
					Subject: "Hello, World!",
					Body:    "<p>This is a test email</p>",
				},
			},
			res: domain.EmailSentEvent{
				From:    "hi@pacheco.io",
				To:      "thiagodelimapacheco@gmail.com",
				Subject: "Hello, World!",
				Body:    "<p>This is a test email</p>",
			},
			err: domain.EmailSendError{
				Message: "Error sending email",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := SendEmail(tt.args.client, tt.args.payload)
			if res != tt.res {
				t.Error("Expected ", tt.res, ", got ", res)
			}
			if (err == nil) != (tt.err == nil) {
				t.Error("Expected ", tt.err, ", got ", err)
			}
		})
	}
}

// Valid email client mock
type EmailClient struct{}

func (client EmailClient) Send(payload *domain.EmailPayload) (domain.EmailSentEvent, error) {
	return domain.EmailSentEvent{
		From:    payload.From,
		To:      payload.To,
		Subject: payload.Subject,
		Body:    payload.Body,
	}, nil
}

// Invalid email client mock
type EmailClientThatFails struct{}

func (client EmailClientThatFails) Send(payload *domain.EmailPayload) (domain.EmailSentEvent, error) {
	res := domain.EmailSentEvent{
		From:    payload.From,
		To:      payload.To,
		Subject: payload.Subject,
		Body:    payload.Body,
	}
	err := domain.EmailSendError{
		Message: "Error sending email",
	}
	return res, err
}
