package activities

import (
	"time"
)

type EmailSender interface {
	Send() error
}

type UserSignUpTimeProvider interface {
	GetUserLastSignUpTime(username string) (time.Time, error)
}

// Activities ...
type Activities struct {
	activityDuration   time.Duration
	signUpTimeProvider UserSignUpTimeProvider
	emailSender        EmailSender
}

// New ...
func New(
	activityDuration time.Duration,
	signUpTimeProvider UserSignUpTimeProvider,
	emailSender EmailSender,
) *Activities {
	return &Activities{
		activityDuration:   activityDuration,
		signUpTimeProvider: signUpTimeProvider,
		emailSender:        emailSender,
	}
}

// Fail ...
func (a Activities) Fail() error {
	time.Sleep(a.activityDuration)
	return nil
}

// Success ...
func (a Activities) Success() error {
	time.Sleep(a.activityDuration)
	return nil
}

// Increase ...
func (a Activities) Increase(counter int) (int, error) {
	time.Sleep(a.activityDuration)
	res := counter + 1
	return res, nil
}

// HasUserSignedUp ...
func (a Activities) HasUserSignedUp(username string) (bool, error) {
	time.Sleep(a.activityDuration)
	t, err := a.signUpTimeProvider.GetUserLastSignUpTime(username)
	if err != nil {
		return false, err
	}

	if !t.IsZero() {
		return true, nil
	}
	return false, nil
}

// SendReminderEmail ...
func (a Activities) SendReminderEmail(user string) error {
	time.Sleep(a.activityDuration)
	return a.emailSender.Send()
}

// SendWelcomeEmail ...
func (a Activities) SendWelcomeEmail(user string) error {
	time.Sleep(a.activityDuration)
	return a.emailSender.Send()
}
