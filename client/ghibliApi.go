package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/common"
	"github.com/camilocorreaUdeA/academy-go-q32021/constants"
	"github.com/camilocorreaUdeA/academy-go-q32021/models"
)

type GhibliApiClient interface {
	GetFilms() ([]models.GhibliFilm, error)
	GetFilmById(id string) (models.GhibliFilm, error)
}

type ghibliClient struct {
	client common.HttpClient
}

// NewGhibliApiClient creates a new instance of the GhibliClient struct
func NewGhibliApiClient(c common.HttpClient) (*ghibliClient, error) {
	if c == nil {
		return &ghibliClient{}, fmt.Errorf("ghibliclient needs an http client to work")
	}
	return &ghibliClient{
		client: c,
	}, nil
}

// GetFilms qeuries the ghibli films API and returns all films
func (gc *ghibliClient) GetFilms() ([]models.GhibliFilm, error) {
	response, err := gc.client.CallApi(http.MethodGet, constants.GhibliApiUrl, constants.FILMS, "")
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
func (gc *ghibliClient) GetFilmById(id string) (models.GhibliFilm, error) {
	response, err := gc.client.CallApi(http.MethodGet, constants.GhibliApiUrl, constants.FILMS, id)
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
