package home

import (
	"github.com/johnazedo/playmovie-bff/src/home/domain"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMovieRepositoryImpl(t *testing.T) {
	var client MockClient
	expected := []domain.Movie{
		{
			ID:       "0001",
			Title:    "Harry Potter",
			ImageURL: "random_image_link",
		},
	}
	repository := MovieRepositoryImpl{Service: &client}

	movies, err := repository.GetMoviesByGenres("01")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, movies)
}

type MockClient struct{}

func (c *MockClient) Get(model any, path string) error {
	rv := reflect.ValueOf(model).Elem()

	newModel := TrailModel{
		Result: []MovieModel{
			{
				ID:         "0001",
				Title:      "Harry Potter",
				PosterPath: "random_image_link",
			},
		},
	}

	rv.Set(reflect.ValueOf(newModel))
	return nil
}
