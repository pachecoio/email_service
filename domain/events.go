package domain

type EmailSentEvent struct {
	From    string
	To      string
	Subject string
	Body    string
	Failed  bool
	Message string
}

type EmailBatchSentEvent struct {
	Emails []EmailSentEvent
}
