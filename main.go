package main

import (
	"log"

	"github.com/LuisFerTR/twitter_clone/db"
	"github.com/LuisFerTR/twitter_clone/handlers"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("No connection to the database")
		return
	}

	handlers.Handlers()
}
