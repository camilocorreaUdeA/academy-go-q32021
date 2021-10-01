package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/client"
	"github.com/camilocorreaUdeA/academy-go-q32021/common"
	"github.com/camilocorreaUdeA/academy-go-q32021/models"
)

// GetFilms handles the request and returns back requested items
// in json encoded response.
func GetFilms(w http.ResponseWriter, r *http.Request) {
	responseFilm := []models.GhibliFilm{}
	requestedFilmID := r.URL.Query().Get("id")
	httpClient := common.NewHttpClient()
	ghibliApiClient, err := client.NewGhibliApiClient(httpClient)
	if err != nil {

	}
	film, err := ghibliApiClient.GetFilmById(requestedFilmID)
	if err != nil {
		log.Printf("Failed to fetch film from api: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	newRecord := filmObjectToRecord(film)
	err = common.UpdateCSVFile(filmsFile, newRecord)
	if err != nil {
		log.Printf("Failed to update CSV file: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	responseFilm = append(responseFilm, film)
	respJson, err := json.Marshal(responseFilm)
	if err != nil {
		log.Printf("Failed to marshal handler response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
	return
}
