package repos

import (
	"fmt"
	"log"

	dbmodels "gitlab.com/alexalreadytaken/notify-service/internal/models/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PgNotifyerRepo struct {
	db *gorm.DB
}

var (
	messagesModel *dbmodels.Message = &dbmodels.Message{}
	clientsModel  *dbmodels.Client  = &dbmodels.Client{}
	mailingsModel *dbmodels.Mailing = &dbmodels.Mailing{}
)

const (
	errSaveClientMsg     = "error while saving client"
	errFindClientInfoMsg = "error while getting info about clients"
	errDeleteClientMsg   = "error while deleting client"

	errSaveMailingMsg     = "error while saving mailing"
	errFindMailingInfoMsg = "error while getting info about mailings"
	errDeleteMailingMsg   = "error while deleting mailing"

	errSaveMessageMsg     = "error while saving message"
	errFindMessageInfoMsg = "error while getting info about messages"
)

func NewPgNotifyerRepo(db *gorm.DB) (*PgNotifyerRepo, error) {
	err := db.AutoMigrate(messagesModel, clientsModel, mailingsModel)
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
		log.Println(errSaveClientMsg, err.Error())
		return 0, fmt.Errorf(errSaveClientMsg)
	}
	return client.ID, nil
}

func (repo *PgNotifyerRepo) GetAllClients() ([]dbmodels.Client, error) {
	var clients []dbmodels.Client
	err := repo.db.
		Model(clientsModel).
		Find(&clients).Error
	if err != nil {
		log.Println(errFindClientInfoMsg, err.Error())
		return nil, fmt.Errorf(errFindClientInfoMsg)
	}
	return clients, nil
}

func (repo *PgNotifyerRepo) ClientExists(id uint) (bool, error) {
	var exists bool
	err := repo.db.Model(clientsModel).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error
	if err != nil {
		log.Println(errFindClientInfoMsg, err.Error())
		return false, fmt.Errorf(errFindClientInfoMsg)
	}
	return exists, err
}

func (repo *PgNotifyerRepo) FindClientById(id uint) (*dbmodels.Client, error) {
	var client dbmodels.Client
	err := repo.db.Find(&client, id).Error
	if err != nil {
		log.Println(errFindClientInfoMsg, err)
		return nil, fmt.Errorf(errFindClientInfoMsg)
	}
	return &client, nil
}

func (repo *PgNotifyerRepo) FindClientsByTag(tag string) ([]dbmodels.Client, error) {
	var clients []dbmodels.Client
	err := repo.db.
		Where("tag = ?", tag).
		Find(&clients).Error
	if err != nil {
		log.Println(errFindClientInfoMsg, err)
		return nil, fmt.Errorf(errFindClientInfoMsg)
	}
	return clients, nil
}

func (repo *PgNotifyerRepo) FindClientsByOperator(operator string) ([]dbmodels.Client, error) {
	var clients []dbmodels.Client
	err := repo.db.
		Where("mobile_operator_code = ?", operator).
		Find(&clients).Error
	if err != nil {
		log.Println(errFindClientInfoMsg, err)
		return nil, fmt.Errorf(errFindClientInfoMsg)
	}
	return clients, nil
}

func (repo *PgNotifyerRepo) UpdateClient(client dbmodels.Client) error {
	err := repo.db.Save(&client).Error
	if err != nil {
		log.Println(errSaveClientMsg, err)
		return fmt.Errorf(errSaveClientMsg)
	}
	return nil
}

func (repo *PgNotifyerRepo) DeleteClient(clientId uint) (*dbmodels.Client, error) {
	var client dbmodels.Client
	err := repo.db.
		Clauses(clause.Returning{}).
		Unscoped(). // delete permanent
		Where("id = ?", clientId).
		Delete(&client).
		Error
	if err != nil {
		log.Println(errDeleteClientMsg, err)
		return nil, fmt.Errorf(errDeleteClientMsg)
	}
	return &client, nil
}

func (repo *PgNotifyerRepo) NewMailing(mailing dbmodels.Mailing) (id uint, err error) {
	mailing.ID = 0
	if err := repo.db.Create(&mailing).Error; err != nil {
		log.Println(errSaveMailingMsg, err.Error())
		return 0, fmt.Errorf(errSaveMailingMsg)
	}
	return mailing.ID, nil
}

func (repo *PgNotifyerRepo) GetAllMailings() ([]dbmodels.Mailing, error) {
	var mailings []dbmodels.Mailing
	err := repo.db.
		Model(mailingsModel).
		Find(&mailings).Error
	if err != nil {
		log.Println(errFindMailingInfoMsg, err.Error())
		return nil, fmt.Errorf(errFindMailingInfoMsg)
	}
	return mailings, nil
}

func (repo *PgNotifyerRepo) MailingExists(id uint) (bool, error) {
	var exists bool
	err := repo.db.Model(mailingsModel).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error
	if err != nil {
		log.Println(errFindMailingInfoMsg, err)
		return false, fmt.Errorf(errFindMailingInfoMsg)
	}
	return exists, nil
}

func (repo *PgNotifyerRepo) FindMailingById(id uint) (*dbmodels.Mailing, error) {
	var mailing dbmodels.Mailing
	err := repo.db.Find(&mailing, id).Error
	if err != nil {
		log.Println(errFindMailingInfoMsg, err)
		return nil, fmt.Errorf(errFindMailingInfoMsg)
	}
	return &mailing, nil
}

func (repo *PgNotifyerRepo) UpdateMailing(mailing dbmodels.Mailing) error {
	err := repo.db.Save(&mailing).Error
	if err != nil {
		log.Println(errSaveMailingMsg, err)
		return fmt.Errorf(errSaveMailingMsg)
	}
	return nil
}

func (repo *PgNotifyerRepo) DeleteMailing(mailingId uint) (*dbmodels.Mailing, error) {
	var client dbmodels.Mailing
	err := repo.db.
		Clauses(clause.Returning{}).
		Unscoped(). // delete permanent
		Where("id = ?", mailingId).
		Delete(&client).
		Error
	if err != nil {
		log.Println(errDeleteMailingMsg, err)
		return nil, fmt.Errorf(errDeleteMailingMsg)
	}
	return &client, nil
}

func (repo *PgNotifyerRepo) NewMessage(message dbmodels.Message) (id uint, err error) {
	message.ID = 0
	if err := repo.db.Create(&message).Error; err != nil {
		log.Println(errSaveMessageMsg, err.Error())
		return 0, fmt.Errorf(errSaveMessageMsg)
	}
	return message.ID, nil
}

func (repo *PgNotifyerRepo) CountMailingMessagesByStatus(mailingId uint, status dbmodels.SendingStatus) (int64, error) {
	var count int64
	err := repo.db.Model(&dbmodels.Message{}).
		Where("mailing_id = ?", mailingId).
		Where("sending_status = ?", status).
		Count(&count).Error
	if err != nil {
		log.Println(errFindMessageInfoMsg, err)
		return 0, fmt.Errorf(errFindMessageInfoMsg)
	}
	return count, nil
}

func (repo *PgNotifyerRepo) GetMailingMessages(mailingId uint) ([]dbmodels.Message, error) {
	var messages []dbmodels.Message
	err := repo.db.Model(&dbmodels.Message{}).
		Where("mailing_id = ?", mailingId).
		Find(&messages).Error
	if err != nil {
		log.Println(errFindMessageInfoMsg, err)
		return nil, fmt.Errorf(errFindMessageInfoMsg)
	}
	return messages, nil
}
