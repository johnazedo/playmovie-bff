package home

import (
	"fmt"
	"github.com/johnazedo/playmovie-bff/src/hermes"
	. "github.com/johnazedo/playmovie-bff/src/home/domain"
)

type MovieRepositoryImpl struct {
	Service hermes.Client
	MovieRepository
}

func (r MovieRepositoryImpl) GetMoviesByGenres(genreId string) ([]Movie, error) {
	var model TrailModel
	url := fmt.Sprintf("/discover/movie?with_genres=%s", genreId)
	err := r.Service.Get(&model, url)
	if err != nil {
		return nil, err
	}
	movies := model.toDomain()
	return movies, nil
}
