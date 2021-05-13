package usecases

import "github.com/carrot-systems/cs-session/src/core/domain"

type Usecases interface {
	CreateSession(request *domain.SessionCreationRequest) (*domain.Session, error)
	FindSession(id string) error
	DeleteSession(id string) error
}

func (i interactor) CreateSession(request *domain.SessionCreationRequest) (*domain.Session, error) {
	return nil, nil
}

func (i interactor) FindSession(id string) error {
	return nil
}

func (i interactor) DeleteSession(id string) error {
	return nil
}
