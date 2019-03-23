package notification

// Email struct
type Email struct {
	Subject   string
	Message   string
	Sender    string
	Recipient []string
}

// Mailer protocol
type Mailer interface {
	Send(message string, to []string) error
}
