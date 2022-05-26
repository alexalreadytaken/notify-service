package rest

import "time"

type Message struct {
	ID                       uint
	ConnectionTimeMillis     int64
	ConnectionIdleTimeMillis int64
	SendingTime              time.Time
	ClientId                 uint
	MailingId                uint
	SendingStatus            string
}
