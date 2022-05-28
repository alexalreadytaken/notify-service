package services

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"gitlab.com/alexalreadytaken/notify-service/internal/models/db"
	"gitlab.com/alexalreadytaken/notify-service/internal/models/rest"
	"gitlab.com/alexalreadytaken/notify-service/internal/repos"
	"gitlab.com/alexalreadytaken/notify-service/internal/utils"
)

type SchedulerService struct {
	notifyerRepo           *repos.PgNotifyerRepo
	externalSenderToken    string
	externalSenderEndpoint string
	resty                  *resty.Client
}

func NewSchedulerService(repo *repos.PgNotifyerRepo, cnf *utils.AppConfig) *SchedulerService {
	return &SchedulerService{
		notifyerRepo:           repo,
		externalSenderToken:    cnf.ExternalSenderToken,
		externalSenderEndpoint: cnf.ExternalSenderEndpoint,
		resty:                  resty.New(),
	}
}

func (service *SchedulerService) ScheduleOrSend(mailingId uint) {
	go service.execute(mailingId)
}

func (service *SchedulerService) execute(mailingId uint) {
	mailing := service.scheduleAndGetMailing(mailingId)
	if mailing == nil {
		return
	}
	service.sendMessagesAndCollectStatistics(mailing)
}

// пользователь может бесконечно отодвигать начало рассылки?
func (service *SchedulerService) scheduleAndGetMailing(mailingId uint) *db.Mailing {
	for {
		mailing, _ := service.notifyerRepo.FindMailingById(mailingId)
		if mailing == nil {
			log.Printf("can`t find mailing by id = %d for start execition", mailingId)
			return nil
		}
		now := time.Now()
		if mailing.StartingTime.After(now) {
			time.Sleep(mailing.StartingTime.Sub(now))
		} else if mailing.EndingTime.Before(now) {
			log.Printf("mailing by id = %d has been edited in the past or ended", mailingId)
			return nil
		} else {
			return mailing
		}
	}
}

func (service *SchedulerService) sendMessagesAndCollectStatistics(mailing *db.Mailing) {
	var clients []db.Client
	var err error
	switch mailing.SendindFilter {
	case db.BY_OPERATOR:
		clients, err = service.notifyerRepo.FindClientsByOperator(mailing.FilterValue)
	case db.BY_TAG:
		clients, err = service.notifyerRepo.FindClientsByTag(mailing.FilterValue)
	default:
		log.Println("unexpected mailing filter found=", mailing.SendindFilter)
		return
	}
	if err != nil {
		log.Printf("can`t get clients with filter %s = %s",
			mailing.SendindFilter, mailing.FilterValue)
		return
	}
	for i := 0; i < len(clients); i++ {
		client := clients[i]
		sendingTime := time.Now()
		resp, err := service.send(&client, mailing)
		ti := resp.Request.TraceInfo()
		msg := db.Message{
			ConnectionTime:     ti.ConnTime,
			ConnectionIdleTime: ti.ConnIdleTime,
			SendingTime:        sendingTime,
			ClientId:           &client.ID,
			MailingId:          &mailing.ID,
		}
		if err != nil {
			log.Println("error while sending message to external service=", err)
			msg.SendingStatus = db.REJECTED
		} else {
			msg.SendingStatus = db.SENT
		}
		service.notifyerRepo.NewMessage(msg)
	}
}

func (service *SchedulerService) send(client *db.Client, mailing *db.Mailing) (*resty.Response, error) {
	return service.resty.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(service.externalSenderToken).
		SetBody(rest.ExternalMessage{
			Id:    client.ID,
			Phone: client.PhoneNumber,
			Text:  mailing.Text,
		}).
		Post(fmt.Sprintf("%s/%d", service.externalSenderEndpoint, client.ID))
}
