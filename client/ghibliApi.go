package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/common"
	"github.com/camilocorreaUdeA/academy-go-q32021/models"
)

type GhibliApiClient interface {
	GetFilms() []models.GhibliFilm
	GetFilmById(id string) (models.GhibliFilm, error)
}

type GhibliClient struct {
	client common.HttpClient
}

func NewGhibliApiClient(c common.HttpClient) (*GhibliClient, error) {
	if c == nil {
		return &GhibliClient{}, fmt.Errorf("ghibliclient needs a http client to work")
	}
	return &GhibliClient{
		client: c,
	}, nil
}

func (gc *GhibliClient) GetFilms() []models.GhibliFilm {
	response, err := gc.client.CallApi(http.MethodGet, ghibliApiUrl, "films", "")
	var films []models.GhibliFilm
	err = json.Unmarshal(response, &films)
	if err != nil {
		fmt.Println(err)
	}
	return films
}

func (gc *GhibliClient) GetFilmById(id string) (models.GhibliFilm, error) {
	response, err := gc.client.CallApi(http.MethodGet, ghibliApiUrl, "films", id)
	if err != nil {
		fmt.Println(err)
		fmt.Println("here")
		return models.GhibliFilm{}, err
	}
	film := models.GhibliFilm{}
	err = json.Unmarshal(response, &film)
	if err != nil {
		fmt.Println(err)
		fmt.Println("here2")
		return models.GhibliFilm{}, err
	}
	return film, nil
}
