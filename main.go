package main

import (
	"log"
	"net/http"
	"os"

	"github.com/camilocorreaUdeA/academy-go-q32021/client"
	"github.com/camilocorreaUdeA/academy-go-q32021/common"
	"github.com/camilocorreaUdeA/academy-go-q32021/constants"
	"github.com/camilocorreaUdeA/academy-go-q32021/handlers"
	"github.com/camilocorreaUdeA/academy-go-q32021/repository"
	"github.com/camilocorreaUdeA/academy-go-q32021/services"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("An error ocurred trying to load env variables: %s", err.Error())
		return
	}

	serviceConfig := ":" + os.Getenv("SERVICE_PORT")
	repo := repository.NewFilmsRepo(constants.FilmsFile)
	client, err := client.NewGhibliApiClient(common.NewHttpClient())
	if err != nil {
		log.Printf("http client construction failed: %s", err)
		return
	}
	service, err := services.NewGhibliService(repo, client)
	if err != nil {
		log.Printf("ghibli service construction failed: %s", err)
		return
	}
	ghibliHandler, err := handlers.NewGhibliHandler(service)
	if err != nil {
		log.Printf("handlers construnction failed: %s", err)
		return
	}

	http.HandleFunc(constants.FilmsRoute, ghibliHandler.GetFilms)
	http.HandleFunc(constants.FilmsRouteParams, ghibliHandler.FilmsMux)
	http.HandleFunc(constants.TestWorkersRoute, ghibliHandler.TestWorkers)
	http.HandleFunc(constants.TestWorkersRoute2, ghibliHandler.GetFilmsConcurrently)
	err = http.ListenAndServe(serviceConfig, nil)
	if err != nil {
		log.Printf("An error ocurred trying to run the service: %s", err.Error())
		return
	}
}
