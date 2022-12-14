package domain

import (
	"errors"
)

type GetHomeUseCase struct {
	MovieRepository
	CatalogRepository
}

func (uc *GetHomeUseCase) Execute() ([]Trail, error) {
	genres, err := uc.CatalogRepository.GetGenresList()
	if err != nil {
		return []Trail{}, err
	}

	var trails []Trail

	for _, genre := range genres {
		movies, err := uc.MovieRepository.GetMoviesByGenres(genre.ID)
		if err != nil {
			return []Trail{}, err
		}
		trails = append(trails, Trail{
			genre.ID,
			genre.Title,
			movies,
		})
	}

	if len(trails) == 0 {
		return nil, errors.New("has no trail")
	}

	return trails, nil
}
