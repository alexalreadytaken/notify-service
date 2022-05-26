package db

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	PhoneNumber        string
	MobileOperatorCode string
	Tag                string
	Timezone           string
}
