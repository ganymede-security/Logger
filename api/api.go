package api

import (
	"context"
	"logger/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type Log struct {
	Id    string `json:"id"`
	Date  string `json:"date"`
	Time  string `json:"time"`
	File  string `json:"file"`
	Level string `json:"level"`
	Msg   string `json:"msg"`
	Args  string `json:"args"`
}

type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

var logsCollection *mongo.Collection = db.GetCollection(db.DbClient, "logs")
var validate = validator.New()

func CreateLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var log Log
		defer cancel()

		if err := c.BindJSON(&log); err != nil {
			c.JSON(http.StatusBadRequest, UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&log); validationErr != nil {
			c.JSON(http.StatusBadRequest, UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newLog := Log{
			Id:    log.Id,
			Date:  log.Date,
			Time:  log.Time,
			File:  log.File,
			Level: log.Level,
			Msg:   log.Msg,
			Args:  log.Args,
		}

		result, err := logsCollection.InsertOne(ctx, newLog)
		if err != nil {
			c.JSON(http.StatusInternalServerError, UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func StartGin() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.POST("logs", CreateLogs())
	router.Run("localhost:8080")
}
