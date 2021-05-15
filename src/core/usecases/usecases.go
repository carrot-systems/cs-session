package usecases

import (
	"github.com/carrot-systems/cs-session/src/core/domain"
)

type Usecases interface {
	CreateSession(request *domain.SessionCreationRequest) (string, error)
	DeleteSession(id string) error
}
