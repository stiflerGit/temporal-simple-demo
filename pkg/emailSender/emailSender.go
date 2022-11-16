package emailSender

import "log"

// EmailSender ...
type EmailSender struct {
}

// Send ...
func (s *EmailSender) Send() error {
	log.Printf("seding email")
	return nil
}
