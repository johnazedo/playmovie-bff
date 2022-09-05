package home

import (
	. "github.com/johnazedo/playmovie-bff/src/home/domain"
	. "github.com/johnazedo/playmovie-bff/src/tools"
)

type MovieRepositoryImpl struct {
	ApiHandler
	MovieRepository
}

func (r MovieRepositoryImpl) GetMovie(id string) (Movie, error) {
	_, err := r.ApiHandler.Get("/movie/" + id)
	if err != nil {
		return Movie{}, err
	}
	return Movie{}, nil
}

func (r MovieRepositoryImpl) GetMoviesByGenres(genreId string) ([]Movie, error) {
	return []Movie{}, nil
}
