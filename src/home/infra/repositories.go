package home

import (
	"fmt"
	"github.com/johnazedo/playmovie-bff/src/hermes"
	. "github.com/johnazedo/playmovie-bff/src/home/domain"
	"log"
)

type MovieRepositoryImpl struct {
	Service hermes.Client
	MovieRepository
}

func (r MovieRepositoryImpl) GetMoviesByGenres(genreId string) ([]Movie, error) {
	var model MoviesModel
	url := fmt.Sprintf("/discover/movie?with_genres=%s", genreId)
	err := r.Service.Get(&model, url)
	if err != nil {
		log.Println("Error on GetMoviesByGenres: " + err.Error())
		return nil, err
	}
	movies := model.toDomain()
	return movies, nil
}

type CatalogRepositoryImpl struct {
	CatalogRepository
}

func (r CatalogRepositoryImpl) GetGenresList() ([]Genre, error) {
	return []Genre{
		{ID: "28", Title: "Action"},
		{ID: "12", Title: "Adventure"},
		{ID: "16", Title: "Animation"},
		{ID: "35", Title: "Comedy"},
		{ID: "80", Title: "Crime"},
		{ID: "99", Title: "Documentary"},
	}, nil
}
