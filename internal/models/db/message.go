package db

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SendingTime        time.Time     `gorm:"column:sending_time"`
	ConnectionTime     time.Duration `gorm:"column:connection_time"`
	ConnectionIdleTime time.Duration `gorm:"column:connection_idle_time"`
	SendingStatus      SendingStatus `gorm:"column:sending_status"`
	ClientId           *uint
	Client             *Client `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MailingId          *uint
	Mailing            *Mailing `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type SendingStatus string

const (
	SENT     SendingStatus = "SENT"
	REJECTED SendingStatus = "REJECTED"
)
