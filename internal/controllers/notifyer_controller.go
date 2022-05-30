package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/alexalreadytaken/notify-service/internal/models/db"
	"gitlab.com/alexalreadytaken/notify-service/internal/models/rest"
	"gitlab.com/alexalreadytaken/notify-service/internal/repos"
	"gitlab.com/alexalreadytaken/notify-service/internal/services"
	"gitlab.com/alexalreadytaken/notify-service/internal/utils"
)

type NotifyerController struct {
	repo      *repos.PgNotifyerRepo
	scheduler *services.SchedulerService
}

func NewNotifyerController(
	repo *repos.PgNotifyerRepo,
	scheduler *services.SchedulerService) *NotifyerController {
	return &NotifyerController{
		repo:      repo,
		scheduler: scheduler,
	}
}

func (c *NotifyerController) NewClient(g *gin.Context) {
	log.Println("new create client request")
	client := bindClient(g)
	if g.IsAborted() {
		return
	}
	id, err := c.repo.NewClient(restClientToDb(client))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	log.Printf("new client crated successfully, id = %d", id)
	g.JSON(http.StatusOK, rest.NewClientResponse{ID: id})
}

func (c *NotifyerController) AllClients(g *gin.Context) {
	log.Println("new request for all client")
	clients, err := c.repo.GetAllClients()
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	g.JSON(http.StatusOK, dbClientsToRest(clients))
}

func (c *NotifyerController) UpdateClient(g *gin.Context) {
	log.Printf("new request for update client")
	client := bindClient(g)
	if g.IsAborted() {
		return
	}
	exists, err := c.repo.ClientExists(client.ID)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	if !exists {
		log.Printf("client with id = %d not found", client.ID)
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "client not found"})
		return
	}
	err = c.repo.UpdateClient(restClientToDb(client))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	log.Printf("client with id = %d updated successfully", client.ID)
	g.JSON(http.StatusOK, rest.Result{Msg: "ok"})

}

func (c *NotifyerController) DeleteClient(g *gin.Context) {
	id := g.Param("id")
	log.Printf("new request for deleting client with id = %s", id)
	clientId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "client id must be number"})
		return
	}
	exists, err := c.repo.ClientExists(uint(clientId))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	if !exists {
		log.Printf("client with id = %d not found", clientId)
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "client not found"})
		return
	}
	client, err := c.repo.DeleteClient(uint(clientId))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	log.Printf("client with id = %d deleted successfully", client.ID)
	g.JSON(http.StatusOK, dbClientToRest(client))
}

func (c *NotifyerController) NewMailing(g *gin.Context) {
	log.Println("new create mailng request")
	mailing := bindMailing(g)
	if g.IsAborted() {
		return
	}
	mailingId, err := c.repo.NewMailing(restMailingToDb(mailing))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	c.scheduler.ScheduleOrSend(mailingId)
	log.Printf("maling with id = %d created successfully", mailingId)
	g.JSON(http.StatusOK, rest.NewMailingResponse{ID: mailingId})
}

func (c *NotifyerController) AllMailings(g *gin.Context) {
	log.Println("request for all mailings")
	mailings, err := c.repo.GetAllMailings()
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	g.JSON(http.StatusOK, dbMailingsToRest(mailings))
}

func (c *NotifyerController) UpdateMailing(g *gin.Context) {
	log.Println("new request for update mailing")
	mailing := bindMailing(g)
	if g.IsAborted() {
		return
	}
	exists, err := c.repo.MailingExists(mailing.ID)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	if !exists {
		log.Printf("mailing with id = %d not found", mailing.ID)
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "mailing not found"})
		return
	}
	err = c.repo.UpdateMailing(restMailingToDb(mailing))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	log.Printf("mailing with id = %d updated successfully", mailing.ID)
	g.JSON(http.StatusOK, rest.Result{Msg: "ok"})
}

func (c *NotifyerController) DeleteMailing(g *gin.Context) {
	id := g.Param("id")
	log.Printf("request for deleting mailign with id = %s", id)
	mailingId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "mailing id must be number"})
		return
	}
	exists, err := c.repo.MailingExists(uint(mailingId))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	if !exists {
		log.Printf("mailign with id = %d not found", mailingId)
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "mailing not found"})
		return
	}
	mailing, err := c.repo.DeleteMailing(uint(mailingId))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	log.Printf("mailing with id = %d deleted successfully", mailingId)
	g.JSON(http.StatusOK, dbMailingToRest(mailing))
}

//kill me
func (c *NotifyerController) MailingsDashboard(g *gin.Context) {
	log.Println("new request for all mailigns dashboard")
	mailings, err := c.repo.GetAllMailings()
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	var dashboard []rest.MailingCountsByStatus
	for i := 0; i < len(mailings); i++ {
		mailing := mailings[i]
		rejectedCount, err := c.repo.CountMailingMessagesByStatus(mailing.ID, db.REJECTED)
		if err != nil {
			g.AbortWithStatusJSON(http.StatusInternalServerError,
				rest.Result{Msg: err.Error()})
			return
		}
		sentCount, err := c.repo.CountMailingMessagesByStatus(mailing.ID, db.SENT)
		if err != nil {
			g.AbortWithStatusJSON(http.StatusInternalServerError,
				rest.Result{Msg: err.Error()})
			return
		}
		var counts []rest.CountMessagesByStatus
		rejected := rest.CountMessagesByStatus{Status: "REJECTED", Count: uint(rejectedCount)}
		sent := rest.CountMessagesByStatus{Status: "SENT", Count: uint(sentCount)}
		counts = append(counts, rejected, sent)
		dashboard = append(dashboard, rest.MailingCountsByStatus{
			Mailing: dbMailingToRest(&mailing),
			Counts:  counts,
		})
	}
	g.JSON(http.StatusOK, rest.MailingsDashboard{
		Dashboard: dashboard,
	})
}

func (c *NotifyerController) MailingStatistics(g *gin.Context) {
	id := g.Param("id")
	log.Printf("new request for statistics for mailing by id = %s", id)
	mailingId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "mailing id must be number"})
		return
	}
	exists, err := c.repo.MailingExists(uint(mailingId))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	if !exists {
		log.Printf("mailing with id = %d not found", mailingId)
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "mailing not found"})
		return
	}
	mailing, err := c.repo.FindMailingById(uint(mailingId))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	messages, err := c.repo.GetMailingMessages(uint(mailingId))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	resp := rest.MailingStatistics{
		Mailing:  dbMailingToRest(mailing),
		Messages: dbMessagesToRest(messages),
	}
	g.JSON(http.StatusOK, resp)
}

func bindClient(g *gin.Context) *rest.Client {
	client := rest.Client{}
	prefix := "error while binding client ="
	if err := g.Bind(&client); err != nil {
		msg := "invalid client format=" + err.Error()
		log.Println(prefix, msg)
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: msg})
		return nil
	}
	if !utils.IsDigit(client.PhoneNumber) {
		msg := "phone number must be a digit"
		log.Println(prefix, msg)
		g.AbortWithStatusJSON(http.StatusBadRequest, rest.Result{Msg: msg})
		return nil
	}
	if strings.Split(client.PhoneNumber, "")[0] != "7" {
		msg := "phone number must starts at 7"
		log.Println(prefix, msg)
		g.AbortWithStatusJSON(http.StatusBadRequest, rest.Result{Msg: msg})
		return nil
	}
	if !utils.IsDigit(client.MobileOperatorCode) {
		msg := "mobile operator code must be a digit"
		log.Println(prefix, msg)
		g.AbortWithStatusJSON(http.StatusBadRequest, rest.Result{Msg: msg})
		return nil
	}
	_, err := time.LoadLocation(client.Timezone)
	if err != nil {
		log.Println(prefix, err.Error())
		g.AbortWithStatusJSON(http.StatusBadRequest, rest.Result{Msg: err.Error()})
		return nil
	}
	return &client
}

func bindMailing(g *gin.Context) *rest.Mailing {
	mailing := rest.Mailing{}
	prefix := "error while binding mailing"
	if err := g.Bind(&mailing); err != nil {
		msg := "invalid mailing format=" + err.Error()
		log.Println(prefix, msg)
		g.AbortWithStatusJSON(http.StatusBadRequest, rest.Result{Msg: msg})
		return nil
	}
	if mailing.EndingTime.Before(mailing.StartingTime) {
		msg := "mailing end time smaller starting time"
		log.Println(prefix, msg)
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: msg})
		return nil
	}
	return &mailing
}
