package services

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/camilocorreaUdeA/academy-go-q32021/client"
	"github.com/camilocorreaUdeA/academy-go-q32021/models"
	"github.com/camilocorreaUdeA/academy-go-q32021/repository"
	"github.com/camilocorreaUdeA/academy-go-q32021/workerspool"
)

type GhibliService interface {
	CreateFilm(query url.Values) error
	GetFilm(query url.Values) (models.GhibliFilm, error)
	GetFilms() ([]models.GhibliFilm, error)
	GetFilmsConcurrently(query url.Values) ([]models.GhibliFilm, error)
}

type ghibliService struct {
	filsmRepo  repository.FilmsRepository
	httpClient client.GhibliApiClient
}

type workerParams struct {
	Type   string
	Record []string
}

// NewGhibliService returns a service instance, used to query ghibli films API and the repository
func NewGhibliService(repo repository.FilmsRepository, client client.GhibliApiClient) (*ghibliService, error) {
	if repo == nil {
		return &ghibliService{}, fmt.Errorf("service requires a repository")
	}
	if client == nil {
		return &ghibliService{}, fmt.Errorf("service requires an http client")
	}
	return &ghibliService{
		filsmRepo:  repo,
		httpClient: client,
	}, nil
}

// GetFilms requests all films in the ghibli API
func (gs *ghibliService) GetFilms() ([]models.GhibliFilm, error) {
	films, err := gs.httpClient.GetFilms()
	if err != nil {
		log.Printf("Failed to fetch films from api: %s", err)
		return []models.GhibliFilm{}, err
	}

	return films, nil
}

// CreateFilm fecthes a film from ghibli API and updates the repository
func (gs *ghibliService) CreateFilm(query url.Values) error {
	requestedFilmID := query.Get("id")

	film, err := gs.httpClient.GetFilmById(requestedFilmID)
	if err != nil {
		log.Printf("Failed to fetch film from api: %s", err)
		return err
	}
	newRecord := filmObjectToRecord(film)
	err = gs.filsmRepo.UpdateCSVFile(newRecord)
	if err != nil {
		log.Printf("Failed to update CSV file: %s", err)
		return err
	}
	return nil
}

// GetFilm retieves a record from the repository
func (gs *ghibliService) GetFilm(query url.Values) (models.GhibliFilm, error) {
	requestedFilmID := query.Get("id")
	films, err := gs.filsmRepo.ReadCSVFile()
	if err != nil {
		log.Printf("Failed to fetch film from repository: %s", err)
		return models.GhibliFilm{}, err
	}
	return filterFilmsById(films, requestedFilmID)
}

// GetFilmsConcurrently reads films from repository using a worker pool
func (gs *ghibliService) GetFilmsConcurrently(query url.Values) ([]models.GhibliFilm, error) {
	typeParam := query.Get("type")
	maxItemsParam := query.Get("items")
	maxItemsPerWorkerParam := query.Get("items_per_workers")

	films, err := gs.filsmRepo.ReadCSVFile()
	if err != nil {
		log.Printf("Failed to fetch film from repository: %s", err)
		return []models.GhibliFilm{}, err
	}

	maxItemsPerWorkerNum, err := strconv.Atoi(maxItemsPerWorkerParam)
	if err != nil {
		log.Println("Failed conversion:", err)
		return []models.GhibliFilm{}, err
	}

	var jobs []*workerspool.Job
	for _, film := range films {
		jobs = append(jobs, workerspool.NewJob(workerProcess, workerParams{
			Type:   typeParam,
			Record: film,
		}))
	}

	maxItemsParamNum, err := strconv.Atoi(maxItemsParam)
	if err != nil {
		log.Println("Failed conversion:", err)
		return []models.GhibliFilm{}, err
	}
	numWorkers := maxItemsParamNum/2 - 1
	if numWorkers <= 0 {
		numWorkers = 1
	}

	pool := workerspool.NewWorkersPool(jobs, numWorkers, maxItemsPerWorkerNum)
	res := pool.Run()
	if res == nil {
		return []models.GhibliFilm{}, nil
	}
	var filmes []models.GhibliFilm

	for _, film := range res {
		filmes = append(filmes, recordToFilmObject(film))
	}

	if len(filmes) > maxItemsParamNum {
		filmes = filmes[:maxItemsParamNum]
	}

	return filmes, nil
}

func filterFilmsById(films [][]string, id string) (models.GhibliFilm, error) {
	for _, film := range films {
		if film[0] == id {
			return recordToFilmObject(film), nil
		}
	}
	return models.GhibliFilm{}, fmt.Errorf("film with id %s not found in repository", id)
}

func workerProcess(params interface{}) []string {
	p := params.(workerParams)
	typeParam := p.Type
	record := p.Record

	isNumber := record[0][0:1] <= "9"

	if typeParam == "even" && isNumber || typeParam == "odd" && !isNumber {
		return record
	}
	return nil
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
