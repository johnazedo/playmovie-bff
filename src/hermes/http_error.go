package hermes

import "errors"

type HttpError struct {
	error
	StatusCode int
}

func NewHttpError(statusCode int) *HttpError {
	return &HttpError{
		errors.New("error in request"), statusCode,
	}
}
