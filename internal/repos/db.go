package repos

import (
	"fmt"

	"gitlab.com/alexalreadytaken/notify-service/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb(cnf *utils.AppConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=notifyer port=%s sslmode=disable",
		cnf.DbHost, cnf.DbUser, cnf.DbPassword, cnf.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error while opening pg connection=%s", err.Error())
	}
	return db, nil
}
