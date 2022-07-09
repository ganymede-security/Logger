package routes

import (
	"logger/controllers"

	"github.com/gin-gonic/gin"
)

func LoggerRoute(router *gin.Engine) {
	// all routes related to logs come here
	router.POST("/log", controllers.CreateLog())

	router.GET("/log/:logId", controllers.GetALog())

	router.DELETE("/log/:logId", controllers.DeleteLog())

	router.GET("/logs", controllers.GetAllLogs())
}
