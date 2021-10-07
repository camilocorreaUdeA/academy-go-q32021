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

type MockGhibliClient struct {
	mock.Mock
}

type MockFilmsRepo struct {
	mock.Mock
}

func (m *MockFilmsRepo) UpdateCSVFile(filename string, record []string) error {
	args := m.Called(filename, record)
	return args.Error(0)
}

func (m *MockFilmsRepo) ReadCSVFile(filename string) ([][]string, error) {
	args := m.Called(filename)
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
	t.Run("Get film succeded", func(t *testing.T) {
		asserter := assert.New(t)
		mockRepo := &MockFilmsRepo{}
		mockClient := &MockGhibliClient{}
		service, err := NewGhibliService(mockRepo, mockClient)
		asserter.Nil(err)
		asserter.NotNil(service)
		film, _ := ioutil.ReadFile("./testdata/film.json")
		filmObj := models.GhibliFilm{}
		_ = json.Unmarshal(film, &filmObj)
		var query = map[string][]string{
			"id": {"1"},
		}
		mockClient.On("GetFilmById", "1").Return(filmObj, nil)
		mockRepo.On("UpdateCSVFile", mock.AnythingOfType("string"), mock.AnythingOfType("[]string")).Return(nil)
		err = service.GetFilm(query)
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
		err = service.GetFilm(query)
		asserter.NotNil(err)
	})
	t.Run("Get film failed -> repo errored", func(t *testing.T) {
		asserter := assert.New(t)
		mockRepo := &MockFilmsRepo{}
		mockClient := &MockGhibliClient{}
		service, err := NewGhibliService(mockRepo, mockClient)
		asserter.Nil(err)
		asserter.NotNil(service)
		film, _ := ioutil.ReadFile("./testdata/film.json")
		filmObj := models.GhibliFilm{}
		_ = json.Unmarshal(film, &filmObj)
		var query = map[string][]string{
			"id": {"1"},
		}
		mockClient.On("GetFilmById", "1").Return(filmObj, nil)
		mockRepo.On("UpdateCSVFile", mock.AnythingOfType("string"), mock.AnythingOfType("[]string")).Return(errors.New("fail!"))
		err = service.GetFilm(query)
		asserter.NotNil(err)
	})
}
