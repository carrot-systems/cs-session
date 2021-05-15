package usecases

import (
	"github.com/carrot-systems/cs-session/src/core/domain"
)

type SessionRepo interface {
	CreateSession(userId string) (*domain.Session, error)
	DeleteSession(sessionId string) error
}

type UserClientGateway interface {
	CheckCredentials(user string, credentials domain.Credentials) (string, error)
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
