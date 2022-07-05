package main

import (
	"logger/api"
	"logger/db"
)

func main() {
	db.CreateDb()
	api.StartGin()
}
