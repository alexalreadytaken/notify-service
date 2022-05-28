package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "gitlab.com/alexalreadytaken/notify-service/docs"
	"gitlab.com/alexalreadytaken/notify-service/internal/controllers"
	"gitlab.com/alexalreadytaken/notify-service/internal/repos"
	"gitlab.com/alexalreadytaken/notify-service/internal/services"
	"gitlab.com/alexalreadytaken/notify-service/internal/utils"
)

// @title Notify service api
// @version 0.1

// @host localhost:2000
// @BasePath /api
// @schemes http
func main() {
	cnf := utils.LoadAppConfigFromEnv()
	pgDb, err := repos.NewDb(cnf)
	if err != nil {
		log.Fatalf("can`t open db connection=%s", err.Error())
	}
	notifyerRepo, err := repos.NewPgNotifyerRepo(pgDb)
	if err != nil {
		log.Fatalf("can`t create notifyer repo=%s", err.Error())
	}
	shceduler := services.NewSchedulerService(notifyerRepo, cnf)
	notifyerController := controllers.NewNotifyerController(notifyerRepo, shceduler)
	router := gin.Default()
	apiGroup := router.Group("/api")
	controllers.AddNotifyerRoutes(apiGroup, notifyerController)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":2000")
}
