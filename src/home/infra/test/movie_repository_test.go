package test

import (
	. "github.com/johnazedo/playmovie-bff/src/home/infra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMoviesByGenres(t *testing.T) {
	var client MockClient
	repository := MovieRepositoryImpl{Service: &client}

	// Given
	genreID := "0001"

	// When
	movies, err := repository.GetMoviesByGenres(genreID)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	assert.Equal(t, ExpectedMoviesEntity, movies)
}
