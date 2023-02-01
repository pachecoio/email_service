package domain

type EmailPayload struct {
	From    string
	To      string
	Subject string
	Body    string
}

type EmailBatchPayload struct {
	Emails []EmailPayload
}
