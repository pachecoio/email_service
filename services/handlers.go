package services

import (
	"github.com/pachecoio/email_service/adapters"
	"github.com/pachecoio/email_service/domain"
)

func SendEmail(
	client adapters.AbstractEmailClient, payload *domain.EmailPayload) error {
	return client.Send(payload)
}

func SendBatch(
	client adapters.AbstractEmailClient, payload *domain.EmailBatchPayload) domain.EmailBatchSentEvent {

	var emails []domain.EmailSentEvent

	for _, email := range payload.Emails {
		err := SendEmail(client, &email)

		res := domain.EmailSentEvent{
			From:    email.From,
			To:      email.To,
			Subject: email.Subject,
			Body:    email.Body,
			Failed:  false,
			Message: "",
		}
		if err != nil {
			res.Failed = true
			res.Message = err.Error()
		}
		emails = append(emails, res)
	}
	return domain.EmailBatchSentEvent{
		Emails: emails,
	}
}
