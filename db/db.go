package db

import (
	"fmt"
	"log"
)

func CreateDb() {
	client, err := InitClient()
	if err != nil {
		log.Printf("Error connecting to Mongo Client")
	}

	collection := client.Database("db1").Collection("logs")
	fmt.Print(collection)
}
