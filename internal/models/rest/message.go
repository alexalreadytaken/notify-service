package rest

import "time"

type Message struct {
	ID          uint
	CreatedAt   time.Time
	SendingTime time.Time
	ClientId    uint
	MailingId   uint
}
