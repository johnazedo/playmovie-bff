package domain

import (
	"errors"
)

type GetHomeUseCase struct {
	MovieRepository
	HomeRepository
}

func (uc *GetHomeUseCase) Execute() ([]Trail, error) {
	genres := uc.HomeRepository.GetGenresList()
	var trails []Trail

	for _, genre := range genres {
		movies, err := uc.MovieRepository.GetMoviesByGenres(genre.ID)
		if err != nil {
			trails = append(trails, Trail{
				genre.ID,
				genre.Title,
				movies,
			})
		}
	}

	if len(trails) > 0 {
		return nil, errors.New("has no trail")
	}

	return trails, nil
}
