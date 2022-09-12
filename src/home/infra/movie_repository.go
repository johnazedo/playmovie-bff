package home

import (
	"github.com/johnazedo/playmovie-bff/src/hermes"
	. "github.com/johnazedo/playmovie-bff/src/home/domain"
)

type MovieRepositoryImpl struct {
	hermes.Client
	MovieRepository
}

func (r MovieRepositoryImpl) GetMovie(id string) (Movie, error) {
	return Movie{}, nil
}

func (r MovieRepositoryImpl) GetMoviesByGenres(genreId string) ([]Movie, error) {
	return []Movie{}, nil
}
