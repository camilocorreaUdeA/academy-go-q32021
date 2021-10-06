package main

import (
	"log"
	"net/http"
	"os"

	"github.com/camilocorreaUdeA/academy-go-q32021/client"
	"github.com/camilocorreaUdeA/academy-go-q32021/common"
	"github.com/camilocorreaUdeA/academy-go-q32021/handlers"
	"github.com/camilocorreaUdeA/academy-go-q32021/repository"
	"github.com/camilocorreaUdeA/academy-go-q32021/services"
	"github.com/joho/godotenv"
)

const (
	itemsRoute = "/items/"
	filmsRoute = "/films/"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("An error ocurred trying to load env variables: %s", err.Error())
		return
	}

	serviceConfig := ":" + os.Getenv("SERVICE_PORT")
	repo := repository.NewFilmsRepo()
	client, err := client.NewGhibliApiClient(common.NewHttpClient())
	if err != nil {

	}
	service, err := services.NewGhibliService(repo, client)
	if err != nil {

	}
	ghibliHandler, err := handlers.NewGhibliHandler(service)
	if err != nil {

	}

	http.HandleFunc(filmsRoute, ghibliHandler.GetFilms)
	http.HandleFunc(filmsRoute, ghibliHandler.GetFilm)
	err = http.ListenAndServe(serviceConfig, nil)
	if err != nil {
		log.Printf("An error ocurred trying to run the service: %s", err.Error())
		return
	}
}
