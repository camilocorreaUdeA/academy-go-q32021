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
	CreateFilm(query url.Values) error
	GetFilm(query url.Values) (models.GhibliFilm, error)
	GetFilms() ([]models.GhibliFilm, error)
}

type GhibliService struct {
	filsmRepo  repository.Repository
	httpClient client.GhibliApiClient
}

// NewGhibliService returns a service instance, used to query ghibli films API and the repository
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

// GetFilms requests all films in the ghibli API
func (gs *GhibliService) GetFilms() ([]models.GhibliFilm, error) {
	films, err := gs.httpClient.GetFilms()
	if err != nil {
		log.Printf("Failed to fetch films from api: %s", err)
		return []models.GhibliFilm{}, err
	}

	return films, nil
}

// CreateFilm fecthes a film from ghibli API and updates the repository
func (gs *GhibliService) CreateFilm(query url.Values) error {
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

// GetFilm retieves a record from the repository
func (gs *GhibliService) GetFilm(query url.Values) (models.GhibliFilm, error) {
	requestedFilmID := query.Get("id")
	films, err := gs.filsmRepo.ReadCSVFile(filmsFile)
	if err != nil {
		log.Printf("Failed to fetch film from repository: %s", err)
		return models.GhibliFilm{}, err
	}
	return filterFilmsById(films, requestedFilmID)
}

func filterFilmsById(films [][]string, id string) (models.GhibliFilm, error) {
	for _, film := range films {
		if film[0] == id {
			return recordToFilmObject(film), nil
		}
	}
	return models.GhibliFilm{}, fmt.Errorf("film with id %s not found in repository", id)
}

func recordToFilmObject(record []string) models.GhibliFilm {
	return models.GhibliFilm{
		ID:                     record[0],
		Title:                  record[1],
		OriginalTitle:          record[2],
		OriginalTitleRomanised: record[3],
		Description:            record[4],
		Director:               record[5],
		Producer:               record[6],
		ReleaseDate:            record[7],
		RunningTime:            record[8],
		RtScore:                record[9],
		People:                 []string{record[10]},
		Species:                []string{record[11]},
		Locations:              []string{record[12]},
		Vehicles:               []string{record[13]},
		Url:                    record[14],
	}
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
