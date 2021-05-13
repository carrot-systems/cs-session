package usecases

import (
	"github.com/carrot-systems/cs-session/src/core/domain"
	"github.com/google/uuid"
)

type SessionRepo interface {
	FindSession(handle string) (*domain.Session, error)
	CreateSession(user domain.Session) error
	DeleteSession(handle string) error
}

type UserClientGateway interface {
	CheckCredentials(user string, credentials domain.Credentials) (uuid.UUID, error)
}

type interactor struct {
	sessionRepo       SessionRepo
	userClientGateway UserClientGateway
}

func NewInteractor(sR SessionRepo, ucG UserClientGateway) interactor {
	return interactor{
		sessionRepo:       sR,
		userClientGateway: ucG,
	}
}
