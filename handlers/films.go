package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/services"
)

type GhibliHandler struct {
	service services.Service
}

func NewGhibliHandler(svc services.Service) (GhibliHandler, error) {
	if svc == nil {
		return GhibliHandler{}, fmt.Errorf("the handler requires a valid service")
	}
	return GhibliHandler{
		service: svc,
	}, nil
}

func (gh GhibliHandler) FilmsMux(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		gh.GetFilms(w, r)
	case http.MethodPost:
		gh.PostFilm(w, r)
	default:
		log.Printf("http method not allowed: %s", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`http method not allowed`))
		return
	}
}

// GetFilm handles the request and returns back requested items
// in json encoded response.
func (gh GhibliHandler) PostFilm(w http.ResponseWriter, r *http.Request) {
	err := gh.service.GetFilm(r.URL.Query())
	if err != nil {
		log.Printf("Failed to marshal handler response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Film was correctly fetched and added to repository (csv file)`))
	return
}

func (gh GhibliHandler) GetFilms(w http.ResponseWriter, r *http.Request) {
	films, err := gh.service.GetFilms()
	if err != nil {
		log.Printf("Failed to fetch films: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonFilms, err := json.Marshal(films)
	if err != nil {
		log.Printf("Failed to marshal handler response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return

	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFilms)
	return
}
