package model

import "time"

// code 1: blocking warnings
// code 2: non blocking warnings
// code 3: users messages
type UserMessage struct {
	Code      int       `json:"code,omitempty"`
	User      string    `json:"user,omitempty"`
	Text      string    `json:"text,omitempty"`
	Timestamp time.Time `json:"time,omitempty"`
}

func (message *UserMessage) Prepare() error { //TODO add an error here
	message.Timestamp = time.Now()

	return nil
}
