package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// checks if Mongo Env variable is correctly loaded and returns it
func EnvMongoURI() string {
	ex, err0 := os.Getwd()
	if err0 != nil {
		panic(err0)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err, ex)
	}

	return os.Getenv("MONGOURI")
}
