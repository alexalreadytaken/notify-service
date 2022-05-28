package controllers

import (
	"github.com/gin-gonic/gin"
)

func AddNotifyerRoutes(rg *gin.RouterGroup, c *NotifyerController) {
	rg.GET("/clients", AllClients(c))
	client := rg.Group("/client")
	{
		client.POST("/", NewClient(c))
		client.PUT("/", UpdateClient(c))
		client.DELETE("/:id", DeleteClient(c))
	}
	rg.GET("/mailings", AllMailings(c))
	mailing := rg.Group("/mailing")
	{
		mailing.POST("/", NewMailing(c))
		mailing.PUT("/", UpdateMailing(c))
		mailing.DELETE("/:id", DeleteMailing(c))
		mailing.GET("/dashboard", MailingsDashboard(c))
		mailing.GET("/:id/statistics", MailingStatistics(c))
	}
}

//↓↓↓костыль чтобы не пачкать код контроллера комментариями↓↓↓

// Notifyer godoc
// @Summary create client
// @Description creating client
// @Schemes
// @Accept json
// @Param client body rest.Client true "client"
// @Produce json
// @Success 200 {object} rest.NewClientResponse
// @Failure 400 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /client/ [post]
func NewClient(c *NotifyerController) func(*gin.Context) {
	return c.NewClient
}

// Notifyer godoc
// @Summary all clients
// @Description return all clients
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {object} []rest.Client
// @Failure 400 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /clients [get]
func AllClients(c *NotifyerController) func(*gin.Context) {
	return c.AllClients
}

// Notifyer godoc
// @Summary update client
// @Description updates client by id
// @Schemes
// @Accept json
// @Param client body rest.Client true "client"
// @Produce json
// @Success 200 {object} rest.Result
// @Failure 400 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /client/ [put]
func UpdateClient(c *NotifyerController) func(*gin.Context) {
	return c.UpdateClient
}

// Notifyer godoc
// @Summary delete client
// @Description deletes client by id
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {object} rest.Client
// @Failure 400 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Param id path int true "client id"
// @Router /client/{id} [delete]
func DeleteClient(c *NotifyerController) func(*gin.Context) {
	return c.DeleteClient
}

// Notifyer godoc
// @Summary create mailing
// @Description creates mailing
// @Schemes
// @Accept json
// @Param client body rest.Mailing true "mailing"
// @Produce json
// @Success 200 {object} rest.NewMailingResponse
// @Failure 400 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /mailing/ [post]
func NewMailing(c *NotifyerController) func(*gin.Context) {
	return c.NewMailing
}

// Notifyer godoc
// @Summary all mailings
// @Description return all mailings
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {object} []rest.Mailing
// @Failure 400 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /mailings [get]
func AllMailings(c *NotifyerController) func(*gin.Context) {
	return c.AllMailings
}

// Notifyer godoc
// @Summary update mailing
// @Description updates mailing by id
// @Schemes
// @Accept json
// @Param client body rest.Mailing true "mailing"
// @Produce json
// @Success 200 {object} rest.Result
// @Failure 400 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /mailing/ [put]
func UpdateMailing(c *NotifyerController) func(*gin.Context) {
	return c.UpdateMailing
}

// Notifyer godoc
// @Summary delete mailing
// @Description deletes mailing by id
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {object} rest.Mailing
// @Failure 400 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Param id path int true "mailing id"
// @Router /mailing/{id} [delete]
func DeleteMailing(c *NotifyerController) func(*gin.Context) {
	return c.DeleteMailing
}

// Notifyer godoc
// @Summary all mailings dashboard
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {object} rest.MailingsDashboard
// @Failure 400 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /mailing/dashboard [get]
func MailingsDashboard(c *NotifyerController) func(*gin.Context) {
	return c.MailingsDashboard
}

// Notifyer godoc
// @Summary detailed mailing statistic
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {object} rest.MailingStatistics
// @Failure 400 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Param id path int true "mailing id"
// @Router /mailing/{id}/statistics [get]
func MailingStatistics(c *NotifyerController) func(*gin.Context) {
	return c.MailingStatistics
}
