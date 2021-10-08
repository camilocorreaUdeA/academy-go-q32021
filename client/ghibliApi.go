package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/common"
	"github.com/camilocorreaUdeA/academy-go-q32021/models"
)

type GhibliApiClient interface {
	GetFilms() ([]models.GhibliFilm, error)
	GetFilmById(id string) (models.GhibliFilm, error)
}

type GhibliClient struct {
	client common.HttpClient
}

// NewGhibliApiClient creates a new instance of the GhibliClient struct
func NewGhibliApiClient(c common.HttpClient) (*GhibliClient, error) {
	if c == nil {
		return &GhibliClient{}, fmt.Errorf("ghibliclient needs an http client to work")
	}
	return &GhibliClient{
		client: c,
	}, nil
}

// GetFilms qeuries the ghibli films API and returns all films
func (gc *GhibliClient) GetFilms() ([]models.GhibliFilm, error) {
	response, err := gc.client.CallApi(http.MethodGet, ghibliApiUrl, "films", "")
	if err != nil {
		log.Printf("api call failed, unable to fetch data: %s", err)
		return []models.GhibliFilm{}, err
	}
	var films []models.GhibliFilm
	err = json.Unmarshal(response, &films)
	if err != nil {
		log.Printf("unable to unmarshal fetched data: %s", err)
	}
	return films, err
}

// GetFilmById queries a single film from ghibli API
func (gc *GhibliClient) GetFilmById(id string) (models.GhibliFilm, error) {
	response, err := gc.client.CallApi(http.MethodGet, ghibliApiUrl, "films", id)
	if err != nil {
		log.Printf("api call failed, unable to fetch data: %s", err)
		return models.GhibliFilm{}, err
	}
	film := models.GhibliFilm{}
	err = json.Unmarshal(response, &film)
	if err != nil {
		log.Printf("unable to unmarshal fetched data: %s", err)
	}
	return film, err
}
