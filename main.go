package main

import (
	"logger/configs"
	"logger/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// run database
	configs.ConnectDB()

	// routes
	routes.LoggerRoute(router)

	router.Run("localhost:6000")
}
