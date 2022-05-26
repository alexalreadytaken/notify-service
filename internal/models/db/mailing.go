package db

import (
	"time"

	"gorm.io/gorm"
)

type Mailing struct {
	gorm.Model
	StartingTime  time.Time     `gorm:"column:starting_time"`
	Text          string        `gorm:"column:text"`
	SendindFilter SendindFilter `gorm:"column:sending_filter"`
	FilterValue   string        `gorm:"column:filter_value"`
	EndingTime    time.Time     `gorm:"column:ending_time"`
}

type SendindFilter string

const (
	BY_TAG      SendindFilter = "BY_TAG"
	BY_OPERATOR SendindFilter = "BY_OPERATOR"
)
