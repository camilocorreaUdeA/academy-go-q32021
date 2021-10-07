package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/camilocorreaUdeA/academy-go-q32021/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetFilm(query url.Values) error {
	args := m.Called(query)
	return args.Error(0)
}

func (m *MockService) GetFilms() ([]models.GhibliFilm, error) {
	args := m.Called()
	return args.Get(0).([]models.GhibliFilm), args.Error(1)
}

func TestGetFilm(t *testing.T) {
	t.Run("Fetch film succeded", func(t *testing.T) {
		asserter := assert.New(t)
		mockService := &MockService{}
		handler, err := NewGhibliHandler(mockService)
		asserter.Nil(err)
		asserter.NotNil(handler)
		mockService.On("GetFilm", mock.AnythingOfType("url.Values")).Return(nil)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/films/", nil)
		handler.GetFilm(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		asserter.Equal(200, resp.StatusCode)
		asserter.Equal("Film was correctly fetched and added to csv file", string(body))
	})
}
