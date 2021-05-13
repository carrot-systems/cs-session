package rest

import (
	"errors"
	"net/http"
)

var (
	ErrFormValidation = errors.New("failed to validate form")
)

func codeForError(err error) int {
	switch err {
	case ErrFormValidation:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
