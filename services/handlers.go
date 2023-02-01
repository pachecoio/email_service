package services

import (
	"github.com/pachecoio/email_service/adapters"
	"github.com/pachecoio/email_service/domain"
)

func SendEmail(
	client adapters.AbstractEmailClient, payload *domain.EmailPayload) domain.EmailSentEvent {
	err := client.Send(payload)

	res := domain.EmailSentEvent{
		From:    payload.From,
		To:      payload.To,
		Subject: payload.Subject,
		Body:    payload.Body,
		Failed:  false,
		Message: "",
	}
	if err != nil {
		res.Failed = true
		res.Message = err.Error()
	}
	return res
}

func SendBatch(
	client adapters.AbstractEmailClient, payload *domain.EmailBatchPayload) domain.EmailBatchSentEvent {

	var emails []domain.EmailSentEvent

	for _, email := range payload.Emails {
		res := SendEmail(client, &email)
		emails = append(emails, res)
	}
	return domain.EmailBatchSentEvent{
		Emails: emails,
	}
}
