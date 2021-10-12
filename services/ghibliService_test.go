package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/camilocorreaUdeA/academy-go-q32021/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const singleFilm = "./testdata/film.json"

type MockGhibliClient struct {
	mock.Mock
}

type MockFilmsRepo struct {
	mock.Mock
}

func (m *MockFilmsRepo) UpdateCSVFile(record []string) error {
	args := m.Called(record)
	return args.Error(0)
}

func (m *MockFilmsRepo) ReadCSVFile() ([][]string, error) {
	args := m.Called()
	return args.Get(0).([][]string), args.Error(1)
}

func (m *MockGhibliClient) GetFilms() ([]models.GhibliFilm, error) {
	args := m.Called()
	return args.Get(0).([]models.GhibliFilm), args.Error(1)
}

func (m *MockGhibliClient) GetFilmById(id string) (models.GhibliFilm, error) {
	args := m.Called(id)
	return args.Get(0).(models.GhibliFilm), args.Error(1)
}

func TestGetFilms(t *testing.T) {
	t.Run("Get all films succeded", func(t *testing.T) {
		asserter := assert.New(t)
		mockRepo := &MockFilmsRepo{}
		mockClient := &MockGhibliClient{}
		service, err := NewGhibliService(mockRepo, mockClient)
		asserter.Nil(err)
		asserter.NotNil(service)
		films := []models.GhibliFilm{
			{
				Title:         "My neighbor Totoro",
				OriginalTitle: "some shit in japanese",
			},
		}
		mockClient.On("GetFilms").Return(films, nil)
		res, err := service.GetFilms()
		asserter.Nil(err)
		asserter.NotEmpty(res)
		asserter.Equal("My neighbor Totoro", res[0].Title)
	})
	t.Run("Get all films failed", func(t *testing.T) {
		asserter := assert.New(t)
		mockRepo := &MockFilmsRepo{}
		mockClient := &MockGhibliClient{}
		service, err := NewGhibliService(mockRepo, mockClient)
		asserter.Nil(err)
		asserter.NotNil(service)
		films := []models.GhibliFilm{}
		mockClient.On("GetFilms").Return(films, errors.New("ouuch!"))
		res, err := service.GetFilms()
		asserter.NotNil(err)
		asserter.Empty(res)
	})

}

func TestGetFilm(t *testing.T) {
	t.Run("Get film from repository succeded", func(t *testing.T) {
		asserter := assert.New(t)
		mockRepo := &MockFilmsRepo{}
		mockClient := &MockGhibliClient{}
		service, err := NewGhibliService(mockRepo, mockClient)
		asserter.Nil(err)
		asserter.NotNil(service)
		films := [][]string{
			{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14"},
		}
		var query = map[string][]string{
			"id": {"0"},
		}
		mockRepo.On("ReadCSVFile").Return(films, nil)
		res, err := service.GetFilm(query)
		asserter.Nil(err)
		asserter.NotEmpty(res)
		asserter.Equal("7", res.ReleaseDate)
	})
	t.Run("Get film from reposirory failed -> reading csv", func(t *testing.T) {
		asserter := assert.New(t)
		mockRepo := &MockFilmsRepo{}
		mockClient := &MockGhibliClient{}
		service, err := NewGhibliService(mockRepo, mockClient)
		asserter.Nil(err)
		asserter.NotNil(service)
		var query = map[string][]string{
			"id": {"0"},
		}
		mockRepo.On("ReadCSVFile").Return([][]string{}, errors.New("ahhh!"))
		res, err := service.GetFilm(query)
		asserter.NotNil(err)
		asserter.Empty(res)
	})

}

func TestCreateFilm(t *testing.T) {
	t.Run("Get film succeded", func(t *testing.T) {
		asserter := assert.New(t)
		mockRepo := &MockFilmsRepo{}
		mockClient := &MockGhibliClient{}
		service, err := NewGhibliService(mockRepo, mockClient)
		asserter.Nil(err)
		asserter.NotNil(service)
		film, _ := ioutil.ReadFile(singleFilm)
		filmObj := models.GhibliFilm{}
		_ = json.Unmarshal(film, &filmObj)
		var query = map[string][]string{
			"id": {"1"},
		}
		mockClient.On("GetFilmById", "1").Return(filmObj, nil)
		mockRepo.On("UpdateCSVFile", mock.AnythingOfType("[]string")).Return(nil)
		err = service.CreateFilm(query)
		asserter.Nil(err)
	})
	t.Run("Get film failed -> client errored", func(t *testing.T) {
		asserter := assert.New(t)
		mockRepo := &MockFilmsRepo{}
		mockClient := &MockGhibliClient{}
		service, err := NewGhibliService(mockRepo, mockClient)
		asserter.Nil(err)
		asserter.NotNil(service)
		filmObj := models.GhibliFilm{}
		var query = map[string][]string{
			"id": {"1"},
		}
		mockClient.On("GetFilmById", "1").Return(filmObj, errors.New("shit!!"))
		err = service.CreateFilm(query)
		asserter.NotNil(err)
	})
	t.Run("Get film failed -> repo errored", func(t *testing.T) {
		asserter := assert.New(t)
		mockRepo := &MockFilmsRepo{}
		mockClient := &MockGhibliClient{}
		service, err := NewGhibliService(mockRepo, mockClient)
		asserter.Nil(err)
		asserter.NotNil(service)
		film, _ := ioutil.ReadFile(singleFilm)
		filmObj := models.GhibliFilm{}
		_ = json.Unmarshal(film, &filmObj)
		var query = map[string][]string{
			"id": {"1"},
		}
		mockClient.On("GetFilmById", "1").Return(filmObj, nil)
		mockRepo.On("UpdateCSVFile", mock.AnythingOfType("[]string")).Return(errors.New("fail!"))
		err = service.CreateFilm(query)
		asserter.NotNil(err)
	})
}
