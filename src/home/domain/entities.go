package domain

type Movie struct {
	ID          string `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageURL    string `json:"image_url" validate:"required"`
}

type Trail struct {
	ID     string  `json:"id" validate:"required"`
	Title  string  `json:"title" validate:"required"`
	Movies []Movie `json:"movies" validate:"required"`
}

type Genre struct {
	ID    string
	Title string
}
