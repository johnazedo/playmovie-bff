package test

import (
	"github.com/johnazedo/playmovie-bff/src/home/domain"
	. "github.com/johnazedo/playmovie-bff/src/home/infra"
	"reflect"
)

var MockMoviesModel = MoviesModel{
	Result: []MovieModel{
		{
			ID:         "0001",
			Title:      "Harry Potter",
			PosterPath: "random_image_link",
		},
	},
}

var ExpectedMoviesEntity = []domain.Movie{
	{
		ID:       "0001",
		Title:    "Harry Potter",
		ImageURL: "random_image_link",
	},
}

type MockClient struct{}

func (c *MockClient) Get(model any, path string) error {
	rv := reflect.ValueOf(model).Elem()
	newModel := MockMoviesModel
	rv.Set(reflect.ValueOf(newModel))
	return nil
}
