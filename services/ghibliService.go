package services

import (
	"fmt"
	"log"
	"net/url"

	"github.com/camilocorreaUdeA/academy-go-q32021/client"
	"github.com/camilocorreaUdeA/academy-go-q32021/models"
	"github.com/camilocorreaUdeA/academy-go-q32021/repository"
)

type Service interface {
	GetFilm(query url.Values) error
	GetFilms() ([]models.GhibliFilm, error)
}

type GhibliService struct {
	filsmRepo  repository.Repository
	httpClient client.GhibliApiClient
}

func NewGhibliService(repo repository.Repository, client client.GhibliApiClient) (*GhibliService, error) {
	if repo == nil {
		return &GhibliService{}, fmt.Errorf("service requires a repository")
	}
	if client == nil {
		return &GhibliService{}, fmt.Errorf("service requires an http client")
	}
	return &GhibliService{
		filsmRepo:  repo,
		httpClient: client,
	}, nil
}

func (gs *GhibliService) GetFilms() ([]models.GhibliFilm, error) {
	films, err := gs.httpClient.GetFilms()
	if err != nil {
		log.Printf("Failed to fetch films from api: %s", err)
		return []models.GhibliFilm{}, err
	}

	return films, nil
}

func (gs *GhibliService) GetFilm(query url.Values) error {
	requestedFilmID := query.Get("id")

	film, err := gs.httpClient.GetFilmById(requestedFilmID)
	if err != nil {
		log.Printf("Failed to fetch film from api: %s", err)
		return err
	}
	newRecord := filmObjectToRecord(film)
	err = gs.filsmRepo.UpdateCSVFile(filmsFile, newRecord)
	if err != nil {
		log.Printf("Failed to update CSV file: %s", err)
		return err
	}
	return nil
}

func filmObjectToRecord(film models.GhibliFilm) []string {
	return []string{
		film.ID,
		film.Title,
		film.OriginalTitle,
		film.OriginalTitleRomanised,
		film.Description,
		film.Director,
		film.Producer,
		film.ReleaseDate,
		film.RunningTime,
		film.RtScore,
		film.People[0],
		film.Species[0],
		film.Locations[0],
		film.Vehicles[0],
		film.Url,
	}
}
