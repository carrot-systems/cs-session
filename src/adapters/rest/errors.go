package rest

import (
	"errors"
	"github.com/carrot-systems/cs-session/src/core/domain"
	"net/http"
)

var (
	ErrFormValidation = errors.New("failed to validate form")
)

func codeForError(err error) int {
	switch err {
	case ErrFormValidation:
		return http.StatusBadRequest
	case domain.ErrUnmarshallingFailed, domain.ErrByteReading:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
