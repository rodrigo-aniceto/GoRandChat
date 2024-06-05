package model

import "time"

type ServerMessage struct {
	Code      int // 1 - ok server, 2 - user message, 3 - error
	User      string
	Text      string
	Timestamp time.Time
}

func (message *ServerMessage) Prepare() error { //TODO add an error here
	return nil
}
