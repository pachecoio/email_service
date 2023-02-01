package main

import (
	"github.com/pachecoio/email_service/adapters/mailgun"
	"github.com/pachecoio/email_service/domain"
	"github.com/pachecoio/email_service/services"
	"os"
)

func main() {

	// Sending hardcoded email as an example:

	err := LoadEnv()
	if err != nil {
		panic(err)
	}
	domainValue := os.Getenv("MAILGUN_DOMAIN")
	apiKeyValue := os.Getenv("MAILGUN_API_KEY")

	//	Send mail with mailgun
	c := mailgun.NewClient(domainValue, apiKeyValue)
	payload := &domain.EmailPayload{
		To:      "thiagodelimapacheco@gmail.com",
		From:    "mailgun@" + domainValue,
		Subject: "Test running app",
		Body:    "Test app running fine",
	}
	_, err = services.SendEmail(c, payload)

	if err != nil {
		panic(err)
	}
}
