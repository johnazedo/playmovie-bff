package home

type TrailResponse struct {
	Result []MovieResponse `json:"results"`
}

type MovieResponse struct {
	ID         string `json:"id"`
	PosterPath string `json:"poster_path"`
}
