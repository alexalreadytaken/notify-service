package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/alexalreadytaken/notify-service/internal/models/rest"
	"gitlab.com/alexalreadytaken/notify-service/internal/repos"
	"gitlab.com/alexalreadytaken/notify-service/internal/utils"
)

type NotifyerController struct {
	repo *repos.PgNotifyerRepo
}

func NewNotifyerController(repo *repos.PgNotifyerRepo) *NotifyerController {
	return &NotifyerController{
		repo: repo,
	}
}

func (c *NotifyerController) NewClient(g *gin.Context) {
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
	g.JSON(http.StatusOK, rest.NewClientResponse{ID: id})

}

func (c *NotifyerController) UpdateClient(g *gin.Context) {
	client := bindClient(g)
	if g.IsAborted() {
		return
	}
	exists, err := c.repo.ClientExists(client.ID)
	if err != nil {
		log.Println("error while getting info about client=", err)
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: "can`t get info about client"})
		return
	}
	if !exists {
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
	g.JSON(http.StatusOK, rest.Result{Msg: "ok"})

}

func (c *NotifyerController) DeleteClient(g *gin.Context) {
	id := g.Param("id")
	clientId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "client id must be number"})
		return
	}
	exists, err := c.repo.ClientExists(uint(clientId))
	if err != nil {
		log.Println("error while getting info about client=", err)
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: "can`t get info about client"})
		return
	}
	if !exists {
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
	g.JSON(http.StatusOK, dbClientToRest(client))
}

func (c *NotifyerController) NewMailing(g *gin.Context) {
	mailing := bindMailing(g)
	if g.IsAborted() {
		return
	}
	id, err := c.repo.NewMailing(restMailingToDb(mailing))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	g.JSON(http.StatusOK, rest.NewMailingResponse{ID: id})

}

func (c *NotifyerController) UpdateMailing(g *gin.Context) {
	mailing := bindMailing(g)
	if g.IsAborted() {
		return
	}
	exists, err := c.repo.MailingExists(mailing.ID)
	if err != nil {
		log.Println("error while getting info about mailing=", err)
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: "can`t get info about mailing"})
		return
	}
	if !exists {
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
	g.JSON(http.StatusOK, rest.Result{Msg: "ok"})

}

func (c *NotifyerController) DeleteMailing(g *gin.Context) {
	id := g.Param("id")
	clientId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "mailing id must be number"})
		return
	}
	exists, err := c.repo.MailingExists(uint(clientId))
	if err != nil {
		log.Println("error while getting info about mailing=", err)
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: "can`t get info about mailing"})
		return
	}
	if !exists {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "mailing not found"})
		return
	}
	mailing, err := c.repo.DeleteMailing(uint(clientId))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			rest.Result{Msg: err.Error()})
		return
	}
	g.JSON(http.StatusOK, dbMailingToRest(mailing))
}

func (c *NotifyerController) MailingsDashboard(g *gin.Context) {

}

func (c *NotifyerController) MailingStatistics(g *gin.Context) {

}

func (c *NotifyerController) NewMailingMessage(g *gin.Context) {

}

func bindClient(g *gin.Context) *rest.Client {
	client := rest.Client{}
	if err := g.Bind(&client); err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "invalid client format=" + err.Error()})
		return nil
	}
	if !utils.IsDigit(client.PhoneNumber) {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "phone number must be a digit"})
		return nil
	}
	if strings.Split(client.PhoneNumber, "")[0] != "7" {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "phone number must starts at 7"})
		return nil
	}
	if !utils.IsDigit(client.MobileOperatorCode) {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "mobile operator code must be a digit"})
		return nil
	}
	_, err := time.LoadLocation(client.Timezone)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: err.Error()})
		return nil
	}
	return &client
}

func bindMailing(g *gin.Context) *rest.Mailing {
	mailing := rest.Mailing{}
	if err := g.Bind(&mailing); err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "invalid mailing format=" + err.Error()})
		return nil
	}
	if mailing.EndingTime.Before(time.Now()) {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "end time of mailing in the past"})
		return nil
	}
	return &mailing
}
