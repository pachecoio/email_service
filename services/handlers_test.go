package services

import (
	"github.com/pachecoio/email_service/adapters"
	"github.com/pachecoio/email_service/domain"
	"reflect"
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
			err: domain.EmailSendError{
				Message: "Error sending email",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SendEmail(tt.args.client, tt.args.payload)
			if (err == nil) != (tt.err == nil) {
				t.Error("Expected ", tt.err, ", got ", err)
			}
		})
	}
}

// Valid email client mock
type EmailClient struct{}

func (client EmailClient) Send(payload *domain.EmailPayload) error {
	return nil
}

// Invalid email client mock
type EmailClientThatFails struct{}

func (client EmailClientThatFails) Send(payload *domain.EmailPayload) error {
	return domain.EmailSendError{
		Message: "Error sending email",
	}
}

func TestSendBatch(t *testing.T) {
	type args struct {
		client  adapters.AbstractEmailClient
		payload *domain.EmailBatchPayload
	}
	tests := []struct {
		name string
		args args
		want domain.EmailBatchSentEvent
	}{
		// Test cases
		{
			name: "Test SendBatch with valid payload and client",
			args: args{
				client: &EmailClient{},
				payload: &domain.EmailBatchPayload{
					Emails: []domain.EmailPayload{
						{
							From:    "hi@pacheco.io",
							To:      "thiagodelimapacheco@gmail.com",
							Subject: "Hello, World!",
							Body:    "<p>This is a test email</p>",
						},
						{
							From:    "hi2@pacheco.io",
							To:      "thiagodelimapacheco@gmail.com",
							Subject: "Hello, World!",
							Body:    "<p>This is a test email</p>",
						},
					},
				},
			},
			want: domain.EmailBatchSentEvent{
				Emails: []domain.EmailSentEvent{
					{
						From:    "hi@pacheco.io",
						To:      "thiagodelimapacheco@gmail.com",
						Subject: "Hello, World!",
						Body:    "<p>This is a test email</p>",
						Failed:  false,
						Message: "",
					},
					{
						From:    "hi2@pacheco.io",
						To:      "thiagodelimapacheco@gmail.com",
						Subject: "Hello, World!",
						Body:    "<p>This is a test email</p>",
						Failed:  false,
						Message: "",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SendBatch(tt.args.client, tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendBatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
