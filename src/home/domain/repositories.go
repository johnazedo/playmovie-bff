package domain

type MovieRepository interface {
	GetMovie(id string) (Movie, error)
	GetMoviesByGenres(genreId string) ([]Movie, error)
}

type HomeRepository interface {
	GetGenresList() []Genre
}
