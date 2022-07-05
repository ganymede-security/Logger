package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Logs struct {
	date  string
	time  string
	file  string
	level string
	msg   string
	args  string
}

var dbClient *mongo.Client
var clientError error
var doOnce sync.Once

const MongoURI = "mongodb://localhost:27017/"

func InitClient() *mongo.Client {
	doOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(MongoURI)

		mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientError = err
		}

		if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
			panic(err)
		}
		dbClient = mongoClient

	})
	return dbClient
}

var DbClient *mongo.Client = InitClient()
