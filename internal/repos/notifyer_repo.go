package repos

import (
	"errors"
	"fmt"
	"log"

	dbmodels "gitlab.com/alexalreadytaken/notify-service/internal/models/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PgNotifyerRepo struct {
	db *gorm.DB
}

func NewPgNotifyerRepo(db *gorm.DB) (*PgNotifyerRepo, error) {
	err := db.AutoMigrate(&dbmodels.Message{}, &dbmodels.Client{}, &dbmodels.Mailing{})
	if err != nil {
		return nil, fmt.Errorf("error while AutoMigrate with notifyer models=%s", err.Error())
	}
	return &PgNotifyerRepo{
		db: db,
	}, nil
}

func (repo *PgNotifyerRepo) NewClient(client dbmodels.Client) (id uint, err error) {
	client.ID = 0
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

func (repo *PgNotifyerRepo) FindClientById(id uint) *dbmodels.Client {
	var client dbmodels.Client
	err := repo.db.Find(&client, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &client
}

func (repo *PgNotifyerRepo) FindClientsByTag(tag string) []dbmodels.Client {
	var clients []dbmodels.Client
	repo.db.Where("tag = ?", tag).Find(&clients)
	return clients
}

func (repo *PgNotifyerRepo) FindClientsByOperator(operator string) []dbmodels.Client {
	var clients []dbmodels.Client
	repo.db.Where("mobile_operator_code = ?", operator).Find(&clients)
	return clients
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
	mailing.ID = 0
	if err := repo.db.Create(&mailing).Error; err != nil {
		msg := "error while saving mailing"
		log.Println(msg, err.Error())
		return 0, fmt.Errorf(msg)
	}
	return mailing.ID, nil
}

func (repo *PgNotifyerRepo) GetAllMailings() []dbmodels.Mailing {
	var mailings []dbmodels.Mailing
	repo.db.Model(&dbmodels.Mailing{}).Find(&mailings)
	return mailings
}

func (repo *PgNotifyerRepo) MailingExists(id uint) (bool, error) {
	var exists bool
	err := repo.db.Model(&dbmodels.Mailing{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error
	return exists, err
}

func (repo *PgNotifyerRepo) FindMailingById(id uint) *dbmodels.Mailing {
	var mailing dbmodels.Mailing
	err := repo.db.Find(&mailing, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &mailing
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

func (repo *PgNotifyerRepo) NewMessage(message dbmodels.Message) (id uint, err error) {
	message.ID = 0
	if err := repo.db.Create(&message).Error; err != nil {
		msg := "error while saving mailing"
		log.Println(msg, err.Error())
		return 0, fmt.Errorf(msg)
	}
	return message.ID, nil
}

func (repo *PgNotifyerRepo) CountMailingMessagesByStatus(mailingId uint, status dbmodels.SendingStatus) int64 {
	var count int64
	err := repo.db.Model(&dbmodels.Message{}).
		Where("mailing_id = ?", mailingId).
		Where("sending_status = ?", status).
		Count(&count).Error
	if err != nil {
		log.Println("error while counting messages=", err)
		return 0
	}
	return count
}

func (repo *PgNotifyerRepo) GetMailingMessages(mailingId uint) ([]dbmodels.Message, error) {
	var messages []dbmodels.Message
	err := repo.db.Model(&dbmodels.Message{}).
		Where("mailing_id = ?", mailingId).
		Find(&messages).Error
	if err != nil {
		msg := "error while getting messages info"
		log.Println(msg, err)
		return nil, fmt.Errorf(msg)
	}
	return messages, nil
}
