package workspace

import (
	"fmt"
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

func (s *Service) Create(userID, title string) (*database.Workspace, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *Service) List(userID string) ([]database.Workspace, error) {

	return nil, fmt.Errorf("not implemented")
}

func (s *Service) Get(workspaceID, userID string) (*database.Workspace, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *Service) CreateInvitationToken(workspaceID, userID string) (string, error) {
	return "", fmt.Errorf("not implemented")
}
