package db

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	PhoneNumber        string `gorm:"column:phone_number"`
	MobileOperatorCode string `gorm:"column:monile_operator_code"`
	Tag                string `gorm:"column:tag"`
	Timezone           string `gorm:"column:timezone"`
}
