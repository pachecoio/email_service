package domain

type EmailSendError struct {
	Message string
	errors  []error
}

func (e EmailSendError) Error() string {
	return e.Message
}
