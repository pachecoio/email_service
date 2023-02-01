package adapters

import "github.com/pachecoio/email_service/domain"

type AbstractEmailClient interface {
	Send(payload *domain.EmailPayload) (domain.EmailSentEvent, error)
}
