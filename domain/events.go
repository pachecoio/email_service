package domain

type EmailSentEvent struct {
	From    string
	To      string
	Subject string
	Body    string
}
