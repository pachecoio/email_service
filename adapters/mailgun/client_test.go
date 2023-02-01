package mailgun

import (
	"github.com/joho/godotenv"
	"github.com/pachecoio/email_service/domain"
	"os"
	"testing"
)

func TestNewMailgunClient(t *testing.T) {
	type args struct {
		domain string
		apiKey string
	}
	tests := []struct {
		name string
		args args
	}{
		// Test cases
		{
			name: "Test NewClient",
			args: args{
				domain: "sandbox123.mailgun.org",
				apiKey: "key-1234567890",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.domain, tt.args.apiKey); got == nil {
				t.Errorf("Failed to create a new Mailgun client")
			}
		})
	}
}

func TestClient_SendFail_InvalidCredentials(t *testing.T) {
	c := NewClient("invalid-domain", "invalid-api-key")
	payload := &domain.EmailPayload{
		To:      "hi@pacheco.io",
		From:    "hi@pacheco.io",
		Subject: "Test",
		Body:    "Test",
	}
	err := c.Send(payload)
	if err == nil {
		t.Errorf("Expected to fail")
	}
}

func TestClient_SendSuccessWithEnvVariables(t *testing.T) {
	// Load environment variables
	err := godotenv.Load("../../.env")

	domainValue := os.Getenv("MAILGUN_DOMAIN")
	apiKeyValue := os.Getenv("MAILGUN_API_KEY")

	if domainValue == "" || apiKeyValue == "" {
		t.Errorf("Missing environment variables")
	}

	c := NewClient(domainValue, apiKeyValue)

	payload := &domain.EmailPayload{
		To:      "thiagodelimapacheco@gmail.com",
		From:    "mailgun@" + domainValue,
		Subject: "Test",
		Body:    "Test",
	}
	err = c.Send(payload)
	if err != nil {
		t.Errorf("Expected to succeed but received error: %v", err)
	}
}
