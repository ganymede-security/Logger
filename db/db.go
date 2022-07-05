package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDb() {
	client := DbClient

	collection := client.Database("db1").Collection("logs")
	fmt.Print(collection)
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("db1").Collection(collectionName)
	return collection
}
