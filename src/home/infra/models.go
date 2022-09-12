package home

import (
	"github.com/johnazedo/playmovie-bff/src/hermes"
	. "github.com/johnazedo/playmovie-bff/src/home/domain"
)

type TrailModel struct {
	hermes.Model
	Result []MovieModel `json:"results"`
}

func (m *TrailModel) toDomain() []Movie {
	var movies []Movie
	for _, model := range m.Result {
		movies = append(movies, model.toDomain())
	}
	return movies
}

type MovieModel struct {
	hermes.Model
	ID         string `json:"id"`
	Title      string `json:"title"`
	PosterPath string `json:"poster_path"`
}

func (m *MovieModel) toDomain() Movie {
	return Movie{
		ID: m.ID, Title: m.Title, ImageURL: m.PosterPath,
	}
}
