package document

import (
	"github.com/yorkie-team/codepair/backend-go/internal/database"
	"github.com/yorkie-team/codepair/backend-go/internal/yorkie"
)

type Service struct {
	db database.Database
	yk *yorkie.Yorkie
}

func NewService(db database.Database, yk *yorkie.Yorkie) *Service {
	return &Service{
		db: db,
		yk: yk,
	}
}

func (s *Service) Create(userID, workspaceID, title string) (*database.Document, error) {
	return nil, nil
}

func (s *Service) UpdateTitle(documentID, userID, title string) error {

	return nil
}

func (s *Service) CreateShareToken(documentID, userID string) (string, error) {
	return "", nil
}

func (s *Service) FindFromSharingToken(tokenString string) (*database.Document, error) {
	return nil, nil
}
