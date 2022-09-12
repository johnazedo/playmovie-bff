package domain

type MovieRepository interface {
	GetMoviesByGenres(genreId string) ([]Movie, error)
}

type CatalogRepository interface {
	GetGenresList() []Genre
}
