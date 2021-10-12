package client

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/camilocorreaUdeA/academy-go-q32021/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	allFilms   = "./testdata/allfilms.json"
	singleFilm = "./testdata/film.json"
)

type MockHttpClient struct {
	mock.Mock
}

func (m *MockHttpClient) CallApi(method, url, resource, id string) ([]byte, error) {
	args := m.Called(method, url, resource, id)
	return args.Get(0).([]byte), args.Error(1)
}

func TestNewGhibliApiClient(t *testing.T) {
	asserter := assert.New(t)
	client := common.NewHttpClient()
	gac, err := NewGhibliApiClient(client)
	asserter.Nil(err)
	asserter.NotNil(gac)
}

func TestGetFilms(t *testing.T) {
	t.Run("Films fetched successfully", func(t *testing.T) {
		asserter := assert.New(t)
		film, _ := ioutil.ReadFile(allFilms)
		mockClient := &MockHttpClient{}
		mockClient.On("CallApi", mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(film, nil)
		gac, err := NewGhibliApiClient(mockClient)
		asserter.Nil(err)
		asserter.NotNil(gac)
		films, err := gac.GetFilms()
		asserter.Nil(err)
		asserter.Equal(21, len(films))
	})
	t.Run("Films fetch failed", func(t *testing.T) {
		asserter := assert.New(t)
		film := []byte{}
		mockClient := &MockHttpClient{}
		mockClient.On("CallApi", mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(film, errors.New("error"))
		gac, err := NewGhibliApiClient(mockClient)
		asserter.Nil(err)
		asserter.NotNil(gac)
		films, err := gac.GetFilms()
		asserter.NotNil(err)
		asserter.Equal(0, len(films))
	})
}

func TestGetFilmById(t *testing.T) {
	t.Run("Film fetched successfully", func(t *testing.T) {
		asserter := assert.New(t)
		film, _ := ioutil.ReadFile(singleFilm)
		mockClient := &MockHttpClient{}
		mockClient.On("CallApi", mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(film, nil)
		gac, err := NewGhibliApiClient(mockClient)
		asserter.Nil(err)
		asserter.NotNil(gac)
		filmObj, err := gac.GetFilmById("5fdfb320-2a02-49a7-94ff-5ca418cae602")
		asserter.Nil(err)
		asserter.NotEmpty(filmObj)
		asserter.Equal("When Marnie Was There", filmObj.Title)
	})

	t.Run("Film fetch failed", func(t *testing.T) {
		asserter := assert.New(t)
		film := []byte{}
		mockClient := &MockHttpClient{}
		mockClient.On("CallApi", mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(film, errors.New("error"))
		gac, err := NewGhibliApiClient(mockClient)
		asserter.Nil(err)
		asserter.NotNil(gac)
		filmObj, err := gac.GetFilmById("5fdfb320-2a02-49a7-94ff-5ca418cae602")
		asserter.NotNil(err)
		asserter.Empty(filmObj)
	})

}
