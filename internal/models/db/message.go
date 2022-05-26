package db

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SendingTime time.Time
	ClientId    *uint
	Client      *Client
	MailingId   *uint
	Mailing     *Mailing
}

type SendingStatus string

const (
	CREATED SendingStatus = "CREATED"
	SENT    SendingStatus = "SENT"
)
