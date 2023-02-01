package integration

import (
	"github.com/joho/godotenv"
	"github.com/pachecoio/email_service/adapters/mailgun"
	"github.com/pachecoio/email_service/domain"
	"github.com/pachecoio/email_service/services"
	"os"
	"testing"
)

func TestSendHandler_Mailgun(t *testing.T) {
	err := godotenv.Load("../../.env")

	if err != nil {
		t.Errorf("Failed to load .env file")
	}
	domainValue := os.Getenv("MAILGUN_DOMAIN")
	apiKeyValue := os.Getenv("MAILGUN_API_KEY")

	mailgunClient := mailgun.NewClient(
		domainValue,
		apiKeyValue,
	)
	if mailgunClient == nil {
		t.Errorf("Failed to create a new Mailgun client")
	}

	payload := &domain.EmailPayload{
		To:      "thiagodelimapacheco@gmail.com",
		From:    "mailgun@" + domainValue,
		Subject: "Test Integration",
		Body:    "Test integration working fine",
	}

	_, err = services.SendEmail(mailgunClient, payload)

	if err != nil {
		t.Errorf("Failed to send email: %s", err.Error())
	}
}
