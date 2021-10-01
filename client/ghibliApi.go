package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/common"
	"github.com/camilocorreaUdeA/academy-go-q32021/models"
)

type GhibliApiClient struct {
	client common.HttpClient
}

func NewGhibliApiClient(c common.HttpClient) (*GhibliApiClient, error) {
	if c == nil {
		return &GhibliApiClient{}, fmt.Errorf("ghibliclient needs a http client to work")
	}
	return &GhibliApiClient{
		client: c,
	}, nil
}

func (gac *GhibliApiClient) GetFilms() []models.GhibliFilm {
	response, err := gac.client.CallApi(http.MethodGet, ghibliApiUrl, "films", "")
	var films []models.GhibliFilm
	err = json.Unmarshal(response, &films)
	if err != nil {
		fmt.Println(err)
	}
	return films
}

func (gac *GhibliApiClient) GetFilmById(id string) (models.GhibliFilm, error) {
	response, err := gac.client.CallApi(http.MethodGet, ghibliApiUrl, "films", id)
	if err != nil {
		fmt.Println(err)
		return models.GhibliFilm{}, err
	}
	film := models.GhibliFilm{}
	err = json.Unmarshal(response, &film)
	if err != nil {
		fmt.Println(err)
		return models.GhibliFilm{}, err
	}
	return film, nil
}
