package home

import (
	. "github.com/johnazedo/playmovie-bff/src/home/domain"
	"strconv"
)

type MoviesModel struct {
	Result []MovieModel `json:"results"`
}

func (m *MoviesModel) toDomain() []Movie {
	var movies []Movie
	for _, model := range m.Result {
		movies = append(movies, model.toDomain())
	}
	return movies
}

type MovieModel struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	PosterPath string `json:"poster_path"`
}

func (m *MovieModel) toDomain() Movie {
	return Movie{
		ID: strconv.FormatInt(m.ID, 10), Title: m.Title, ImageURL: "https://image.tmdb.org/t/p/w500" + m.PosterPath,
	}
}
