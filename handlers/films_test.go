package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/camilocorreaUdeA/academy-go-q32021/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const allFilms = "./testdata/allfilms.json"

type MockService struct {
	mock.Mock
}

func (m *MockService) CreateFilm(query url.Values) error {
	args := m.Called(query)
	return args.Error(0)
}

func (m *MockService) GetFilm(query url.Values) (models.GhibliFilm, error) {
	args := m.Called(query)
	return args.Get(0).(models.GhibliFilm), args.Error(1)
}

func (m *MockService) GetFilms() ([]models.GhibliFilm, error) {
	args := m.Called()
	return args.Get(0).([]models.GhibliFilm), args.Error(1)
}

func TestFilmsMux(t *testing.T) {
	t.Run("Get request", func(t *testing.T) {
		asserter := assert.New(t)
		mockService := &MockService{}
		handler, err := NewGhibliHandler(mockService)
		asserter.Nil(err)
		asserter.NotNil(handler)
		film := models.GhibliFilm{
			ID:            "1",
			Title:         "Something",
			OriginalTitle: "Something in japanese",
		}
		mockService.On("GetFilm", mock.AnythingOfType("url.Values")).Return(film, nil)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/films/", nil)
		handler.FilmsMux(w, r)
		resp := w.Result()
		asserter.Equal(200, resp.StatusCode)
	})

	t.Run("Post request", func(t *testing.T) {
		asserter := assert.New(t)
		mockService := &MockService{}
		handler, err := NewGhibliHandler(mockService)
		asserter.Nil(err)
		asserter.NotNil(handler)
		mockService.On("CreateFilm", mock.AnythingOfType("url.Values")).Return(nil)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/films/", nil)
		handler.FilmsMux(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		asserter.Equal(200, resp.StatusCode)
		asserter.Equal("{\"response\":\"Film was correctly fetched and added to repository (csv file)\"}\n", string(body))
	})

	t.Run("Other request", func(t *testing.T) {
		asserter := assert.New(t)
		mockService := &MockService{}
		handler, err := NewGhibliHandler(mockService)
		asserter.Nil(err)
		asserter.NotNil(handler)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPatch, "/films/", nil)
		handler.FilmsMux(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		asserter.Equal(400, resp.StatusCode)
		asserter.Equal("{\"response\":\"http method not allowed\"}\n", string(body))
	})
}

func TestPostFilm(t *testing.T) {
	t.Run("Fetch and update single film succeded", func(t *testing.T) {
		asserter := assert.New(t)
		mockService := &MockService{}
		handler, err := NewGhibliHandler(mockService)
		asserter.Nil(err)
		asserter.NotNil(handler)
		mockService.On("CreateFilm", mock.AnythingOfType("url.Values")).Return(nil)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/films/", nil)
		handler.PostFilm(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		asserter.Equal(200, resp.StatusCode)
		asserter.Equal("{\"response\":\"Film was correctly fetched and added to repository (csv file)\"}\n", string(body))
	})
}

func TestGetFilm(t *testing.T) {
	t.Run("Fetch single film succeded", func(t *testing.T) {
		asserter := assert.New(t)
		mockService := &MockService{}
		handler, err := NewGhibliHandler(mockService)
		asserter.Nil(err)
		asserter.NotNil(handler)
		film := models.GhibliFilm{
			ID:            "1",
			Title:         "Something",
			OriginalTitle: "Something in japanese",
		}
		mockService.On("GetFilm", mock.AnythingOfType("url.Values")).Return(film, nil)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/films/", nil)
		handler.GetFilm(w, r)
		resp := w.Result()
		asserter.Equal(200, resp.StatusCode)
	})
}

func TestGetFilms(t *testing.T) {
	films, _ := ioutil.ReadFile(allFilms)
	t.Run("Fetch all films succeded", func(t *testing.T) {
		asserter := assert.New(t)
		mockService := &MockService{}
		handler, err := NewGhibliHandler(mockService)
		asserter.Nil(err)
		asserter.NotNil(handler)
		var filmes []models.GhibliFilm
		_ = json.Unmarshal(films, &filmes)
		mockService.On("GetFilms").Return(filmes, nil)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/films", nil)
		handler.GetFilms(w, r)
		resp := w.Result()
		asserter.Equal(200, resp.StatusCode)
	})
}
