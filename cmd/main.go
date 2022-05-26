package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "gitlab.com/alexalreadytaken/notify-service/docs"
	"gitlab.com/alexalreadytaken/notify-service/internal/controllers"
	"gitlab.com/alexalreadytaken/notify-service/internal/repos"
	"gitlab.com/alexalreadytaken/notify-service/internal/routes"
	"gitlab.com/alexalreadytaken/notify-service/internal/utils"
)

// @title Notify service api
// @version 0.1

// @host localhost:2000
// @BasePath /api
// @schemes http
func main() {
	cnf := utils.LoadAppConfigFromEnv()
	notifyerRepo, err := repos.NewPgNotifyerRepo(cnf)
	if err != nil {
		log.Fatalf("can`t create notifyer repo=%s", err.Error())
	}
	notifyerController := controllers.NewNotifyerController(notifyerRepo)
	router := gin.Default()
	apiGroup := router.Group("/api")
	routes.AddNotifyerRoutes(apiGroup, notifyerController)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":2000")
}