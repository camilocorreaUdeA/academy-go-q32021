package handlers

import (
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/client"
	"github.com/camilocorreaUdeA/academy-go-q32021/common"
	"github.com/camilocorreaUdeA/academy-go-q32021/repository"
	"github.com/camilocorreaUdeA/academy-go-q32021/services"
)

// GetFilms handles the request and returns back requested items
// in json encoded response.
func GetFilms(w http.ResponseWriter, r *http.Request) {
	client, err := client.NewGhibliApiClient(common.NewHttpClient())
	if err != nil {
		log.Printf("Failed to marshal handler response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	repo := &repository.FilmsRepository{}
	svc, err := services.NewGhibliService(repo, client)
	svc.GetFilm(r.URL.Query())

	if err != nil {
		log.Printf("Failed to marshal handler response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Film was correctly fetched and added to csv file`))
	return
}
