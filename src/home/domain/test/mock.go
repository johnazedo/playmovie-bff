package test

import (
	"github.com/golang/mock/gomock"
	. "github.com/johnazedo/playmovie-bff/src/home/domain"
)

type MockMovieRepository struct {
	ctrl *gomock.Controller
}

func NewMockMovieRepository(ctrl *gomock.Controller) *MockMovieRepository {
	mock := &MockMovieRepository{ctrl}
	return mock
}

func (m *MockMovieRepository) GetMoviesByGenres(genreId string) ([]Movie, error) {
	return []Movie{}, nil
}
