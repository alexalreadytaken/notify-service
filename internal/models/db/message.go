package db

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SendingTime        time.Time
	ConnectionTime     time.Duration
	ConnectionIdleTime time.Duration
	SendingStatus      SendingStatus
	ClientId           *uint
	Client             *Client
	MailingId          *uint
	Mailing            *Mailing
}

type SendingStatus string

const (
	SENT     SendingStatus = "SENT"
	REJECTED SendingStatus = "REJECTED"
)
