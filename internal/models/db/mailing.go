package db

import (
	"time"

	"gorm.io/gorm"
)

type Mailing struct {
	gorm.Model
	StartingTime  time.Time
	Text          string
	SendindFilter SendindFilter
	FilterValue   string
	EndingTime    time.Time
}

type SendindFilter string

const (
	BY_TAG      SendindFilter = "BY_TAG"
	BY_OPERATOR SendindFilter = "BY_OPERATOR"
)
