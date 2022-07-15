package main

import (
	"logger/configs"
	"logger/interfaces"
	"logger/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	var standardLogger = interfaces.NewLogger()

	var args = "12314"

	standardLogger.InvalidArgValue("client", args)

	// run database
	configs.ConnectDB()

	// routes
	routes.LoggerRoute(router)

	router.Run("localhost:6000")
}
