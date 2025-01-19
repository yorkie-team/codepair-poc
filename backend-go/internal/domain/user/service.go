package user

import (
	"errors"
	"github.com/yorkie-team/codepair/backend-go/internal/database"
)

type Service struct {
	db database.Database
}

func NewService(db database.Database) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) FindByID(userID string) error {
	return errors.New("not implemented")
}

func (s *Service) FindBySocialID(provider, socialUID string) error {
	return errors.New("not implemented")

}

func (s *Service) Create(provider, socialUID, nickname string) error {
	return errors.New("not implemented")
}

func (s *Service) ChangeNickname(userID, nickname string) error {
	return errors.New("not implemented")
}
