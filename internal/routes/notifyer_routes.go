package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddNotifyerRoutes(rg *gin.RouterGroup) {
	rg.GET("/test", handle)
}

// Notifyer godoc
// @Summary test summary
// @Description test description
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /test [get]
func handle(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}
