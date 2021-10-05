package services

import (
	"fmt"
	"log"
	"net/url"

	"github.com/camilocorreaUdeA/academy-go-q32021/client"
	"github.com/camilocorreaUdeA/academy-go-q32021/repository"
)

type Service interface {
	GetFilm(query url.Values) error
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
