package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/services"
)

type GhibliHandler interface {
	FilmsMux(w http.ResponseWriter, r *http.Request)
	PostFilm(w http.ResponseWriter, r *http.Request)
	GetFilm(w http.ResponseWriter, r *http.Request)
	GetFilms(w http.ResponseWriter, r *http.Request)
}
type ghibliHandler struct {
	service services.GhibliService
}

// NewGhibliHandler returns a handler with the functions that can be attached to the endpoints
func NewGhibliHandler(svc services.GhibliService) (ghibliHandler, error) {
	if svc == nil {
		return ghibliHandler{}, fmt.Errorf("the handler requires a valid service")
	}
	return ghibliHandler{
		service: svc,
	}, nil
}

// FilmsMux multiplexes different requests made to the same endpoint "/films/"
func (gh ghibliHandler) FilmsMux(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		gh.GetFilm(w, r)
	case http.MethodPost:
		gh.PostFilm(w, r)
	default:
		log.Printf("http method not allowed: %s", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`http method not allowed`))
		return
	}
}

// PostFilm queries the ghibli API and adds the film to the repository.
func (gh ghibliHandler) PostFilm(w http.ResponseWriter, r *http.Request) {
	err := gh.service.CreateFilm(r.URL.Query())
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

// GetFilm fetchs a film currently stored in repository
func (gh ghibliHandler) GetFilm(w http.ResponseWriter, r *http.Request) {
	film, err := gh.service.GetFilm(r.URL.Query())
	if err != nil {
		log.Printf("Failed to fetch film: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonFilm, err := json.Marshal(film)
	if err != nil {
		log.Printf("Failed to marshal handler response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFilm)
	return
}

// GetFilms retrieves all films in the ghibli films API
func (gh ghibliHandler) GetFilms(w http.ResponseWriter, r *http.Request) {
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
