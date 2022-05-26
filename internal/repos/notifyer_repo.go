package repos

import (
	"fmt"
	"log"

	dbmodels "gitlab.com/alexalreadytaken/notify-service/internal/models/db"
	"gitlab.com/alexalreadytaken/notify-service/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PgNotifyerRepo struct {
	db *gorm.DB
}

func NewPgNotifyerRepo(cnf *utils.AppConfig) (*PgNotifyerRepo, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=notifyer port=%s sslmode=disable",
		cnf.DbHost, cnf.DbUser, cnf.DbPassword, cnf.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error while opening pg connection=%s", err.Error())
	}
	err = db.AutoMigrate(&dbmodels.Message{}, &dbmodels.Client{}, &dbmodels.Mailing{})
	if err != nil {
		return nil, fmt.Errorf("error while AutoMigrate with notifyer models=%s", err.Error())
	}
	return &PgNotifyerRepo{
		db: db,
	}, nil
}

func (repo *PgNotifyerRepo) NewClient(client dbmodels.Client) (id uint, err error) {
	client.ID=0
	if err := repo.db.Create(&client).Error; err != nil {
		msg := "error while saving client"
		log.Println(msg, err.Error())
		return 0, fmt.Errorf(msg)
	}
	return client.ID, nil
}

func (repo *PgNotifyerRepo) ClientExists(id uint) (bool, error) {
	var exists bool
	err := repo.db.Model(&dbmodels.Client{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error
	return exists, err
}

func (repo *PgNotifyerRepo) UpdateClient(client dbmodels.Client) error {
	return repo.db.Save(&client).Error
}

func (repo *PgNotifyerRepo) DeleteClient(clientId uint) (*dbmodels.Client, error) {
	var client dbmodels.Client
	err := repo.db.
		Clauses(clause.Returning{}).
		Unscoped().
		Where("id = ?", clientId).
		Delete(&client).
		Error
	if err != nil {
		msg := "error while deleting client"
		log.Println(msg, err)
		return nil, fmt.Errorf(msg)
	}
	return &client, nil
}

func (repo *PgNotifyerRepo) NewMailing(mailing dbmodels.Mailing) (id uint, err error) {
	mailing.ID=0
	if err := repo.db.Create(&mailing).Error; err != nil {
		msg := "error while saving mailing"
		log.Println(msg, err.Error())
		return 0, fmt.Errorf(msg)
	}
	return mailing.ID, nil
}

//todo to generic?
func (repo *PgNotifyerRepo) MailingExists(id uint) (bool, error) {
	var exists bool
	err := repo.db.Model(&dbmodels.Mailing{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error
	return exists, err
}

func (repo *PgNotifyerRepo) UpdateMailing(mailing dbmodels.Mailing) error {
	return repo.db.Save(&mailing).Error
}

func (repo *PgNotifyerRepo) DeleteMailing(mailingId uint) (*dbmodels.Mailing, error) {
	var client dbmodels.Mailing
	err := repo.db.
		Clauses(clause.Returning{}).
		Unscoped().
		Where("id = ?", mailingId).
		Delete(&client).
		Error
	if err != nil {
		msg := "error while deleting mailing"
		log.Println(msg, err)
		return nil, fmt.Errorf(msg)
	}
	return &client, nil
}
