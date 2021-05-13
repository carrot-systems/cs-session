package gateway

import (
	"github.com/carrot-systems/cs-session/src/core/domain"
	"github.com/carrot-systems/cs-session/src/core/usecases"
	"github.com/google/uuid"
)

type userClientGateway struct {
}

func (u userClientGateway) CheckCredentials(user string, credentials domain.Credentials) (uuid.UUID, error) {
	panic("implement me")
}

func NewUserClientGateway() usecases.UserClientGateway {
	return &userClientGateway{}
}
