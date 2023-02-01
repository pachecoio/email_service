package domain

type EmailSentEvent struct {
	From    string
	To      string
	Subject string
	Body    string
}

type EmailBatchSentEvent struct {
	emails []EmailSentEvent
}
