package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type logs struct {
	id    string `json:"id"`
	date  string `json:"date"`
	time  string `json:"time"`
	file  string `json:"file"`
	level string `json:"level"`
	msg   string `json:"msg"`
	args  string `json:"args"`
}

func getLogs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, logs)
}

func StartGin() {
	router := gin.Default()
	router.GET(logs, getLogs)

	router.Run("localhost:8080")
}
