package controllers

import (
	"context"
	"logger/configs"
	"logger/models"
	"logger/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var logsCollection *mongo.Collection = configs.GetCollection(configs.DB, "logs")
var validate = validator.New()

func CreateLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var log models.Log
		defer cancel()

		// validate request body
		if err := c.BindJSON(&log); err != nil {
			c.JSON(http.StatusBadRequest, responses.LoggerResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// validate required fields
		if validationErr := validate.Struct(&log); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LoggerResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newLog := models.Log{
			Id:    primitive.NewObjectID(),
			Date:  log.Date,
			Time:  log.Time,
			File:  log.File,
			Level: log.Level,
			Msg:   log.Msg,
			Args:  log.Args,
		}

		// insert the result to the db, or an error if unable to
		result, err := logsCollection.InsertOne(ctx, newLog)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LoggerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.LoggerResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetALog() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		logId := c.Param("logId")
		var log models.Log
		defer cancel()

		// search for log with matching Id
		objId, _ := primitive.ObjectIDFromHex(logId)

		// return error if unable to find log, and the log if found
		err := logsCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&log)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LoggerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.LoggerResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": log}})
	}
}

func DeleteLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		logId := c.Param("logId")
		defer cancel()

		// search for log by Id
		objId, _ := primitive.ObjectIDFromHex(logId)

		// return error if unable to find log, and delete log if found
		result, err := logsCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LoggerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.LoggerResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}})
		}

		c.JSON(http.StatusOK, responses.LoggerResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "log successfully deleted"}})
	}
}

func GetAllLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var logs []models.Log
		defer cancel()

		results, err := logsCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LoggerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Read from db
		defer results.Close(ctx)

		for results.Next(ctx) {
			var singleLog models.Log
			if err = results.Decode(&singleLog); err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoggerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			logs = append(logs, singleLog)
		}

		c.JSON(http.StatusOK, responses.LoggerResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": logs}})
	}
}
