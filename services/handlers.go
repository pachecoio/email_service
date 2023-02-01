package services

import (
	"github.com/pachecoio/email_service/adapters"
	"github.com/pachecoio/email_service/domain"
)

func SendEmail(
	client adapters.AbstractEmailClient, payload *domain.EmailPayload) (domain.EmailSentEvent, error) {
	return client.Send(payload)
}
