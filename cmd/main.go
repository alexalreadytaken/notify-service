package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gitlab.com/alexalreadytaken/notify-service/internal/routes"
	_ "gitlab.com/alexalreadytaken/notify-service/docs"
)

// @title Notify service api
// @version 0.1

// @host localhost:2000
// @BasePath /api
// @schemes http
func main() {
	router := gin.Default()

	apiGroup := router.Group("/api")
	routes.AddNotifyerRoutes(apiGroup)
	
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":2000")
}
