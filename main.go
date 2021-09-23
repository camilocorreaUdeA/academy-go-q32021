package main

import (
	"log"
	"net/http"
	"os"

	"github.com/camilocorreaUdeA/academy-go-q32021/handlers"
	"github.com/joho/godotenv"
)

const itemsRoute = "/items/"

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("An error ocurred trying to load env variables: %s", err.Error())
		return
	}

	serviceConfig := ":" + os.Getenv("SERVICE_PORT")

	http.HandleFunc(itemsRoute, handlers.GetItems)
	err = http.ListenAndServe(serviceConfig, nil)
	if err != nil {
		log.Printf("An error ocurred trying to run the service: %s", err.Error())
		return
	}
}
