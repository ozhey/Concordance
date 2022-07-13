package main

import (
	"github.com/joho/godotenv"
	"github.com/ozhey/concordance/controller/api"
	"github.com/ozhey/concordance/controller/database"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	err = database.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	router := api.SetupRouter()
	err = router.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}
}
