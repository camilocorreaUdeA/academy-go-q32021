package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/common"
	"github.com/camilocorreaUdeA/academy-go-q32021/models"
)

const ghibliApiUrl = "ghibliapi.herokuapp.com"

func GetFilms(url string) []*models.GhibliFilm {
	response, err := common.CallApi(http.MethodGet, ghibliApiUrl, "films", "")
	var films []*models.GhibliFilm
	err = json.Unmarshal(response, &films)
	if err != nil {
		fmt.Println(err)
	}
	return films
}

func GetFilmById(url, id string) *models.GhibliFilm {
	response, err := common.CallApi(http.MethodGet, ghibliApiUrl, "films", id)
	film := &models.GhibliFilm{}
	err = json.Unmarshal(response, film)
	if err != nil {
		fmt.Println(err)
	}
	return film
}
