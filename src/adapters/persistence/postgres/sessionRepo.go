package postgres

import (
	"github.com/carrot-systems/cs-session/src/core/domain"
	"github.com/carrot-systems/cs-session/src/core/usecases"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID          string `gorm:"type:uuid;primary_key"`
	Handle      string
	DisplayName string
	Mail        string
	Password    string
}

type sessionRepo struct {
	db *gorm.DB
}

func (s sessionRepo) FindSession(handle string) (*domain.Session, error) {
	panic("implement me")
}

func (s sessionRepo) CreateSession(user domain.Session) error {
	panic("implement me")
}

func (s sessionRepo) DeleteSession(handle string) error {
	panic("implement me")
}

func NewSessionRepo(db *gorm.DB) usecases.SessionRepo {
	return &sessionRepo{
		db,
	}
}
