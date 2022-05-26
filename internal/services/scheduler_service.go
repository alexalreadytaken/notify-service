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
	mailing := service.notifyerRepo.FindMailingById(mailingId)
	if mailing == nil {
		log.Printf("can`t find mailing by id = %d for start execition", mailingId)
		return
	}
	now := time.Now()
	if mailing.StartingTime.Before(now) && mailing.EndingTime.After(now) {
		go service.sendMessagesAndCollectStatistics(mailing)
		return
	}
	go func() {
		time.Sleep(mailing.StartingTime.Sub(now))
		mailing := service.notifyerRepo.FindMailingById(mailingId)
		if mailing == nil {
			log.Printf("can`t find mailing by id = %d for start execition", mailingId)
			return
		}
		service.sendMessagesAndCollectStatistics(mailing)
	}()
}

func (service *SchedulerService) sendMessagesAndCollectStatistics(mailing *db.Mailing) {
	var clients []db.Client
	switch mailing.SendindFilter {
	case db.BY_OPERATOR:
		clients = service.notifyerRepo.FindClientsByOperator(mailing.FilterValue)
	case db.BY_TAG:
		clients = service.notifyerRepo.FindClientsByTag(mailing.FilterValue)
	default:
		log.Println("unexpected mailing filter found=", mailing.SendindFilter)
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
