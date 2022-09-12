package hermes

import "errors"

type HttpError struct {
	error
	StatusCode int
}

func NewHttpError(message string, statusCode int) *HttpError {
	return &HttpError{
		errors.New(message), statusCode,
	}
}
