package postgres

import (
	"github.com/carrot-systems/cs-session/src/core/domain"
	"github.com/carrot-systems/cs-session/src/core/usecases"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID     string `gorm:"type:uuid;primary_key"`
	UserId string
}

type sessionRepo struct {
	db *gorm.DB
}

func (s sessionRepo) CreateSession(userId string) (*domain.Session, error) {
	session := Session{
		Model:  gorm.Model{},
		ID:     uuid.New().String(),
		UserId: userId,
	}

	query := s.db.Create(&session)

	err := query.Error

	if err != nil {
		return nil, err
	}

	return session.toDomain(), nil
}

func (s sessionRepo) DeleteSession(sessionId string) error {
	panic("implement me")
}

func (s Session) toDomain() *domain.Session {
	return &domain.Session{
		ID:     s.ID,
		UserId: s.UserId,
	}
}

func NewSessionRepo(db *gorm.DB) usecases.SessionRepo {
	return &sessionRepo{
		db,
	}
}
