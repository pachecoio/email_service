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

	for _, email := range payload.Emails {
		err := SendEmail(client, &email)
		if err != nil {
			//..
		}
	}
	return domain.EmailBatchSentEvent{}
}
