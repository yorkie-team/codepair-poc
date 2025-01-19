package auth

import (
	"context"
	"fmt"
	"github.com/yorkie-team/codepair/backend-go/internal/database"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/auth/github"
	"github.com/yorkie-team/codepair/backend-go/internal/token"
)

type Service struct {
	github       *github.OAuth
	tokenManager token.Manager
	database     database.Database
	config       *Config
}

func NewService(config *Config, db database.Database, tm *token.Manager) *Service {
	g := github.New(config.Github)
	return &Service{
		database:     db,
		tokenManager: *tm,
		github:       g,
		config:       config,
	}
}

func (s *Service) GetFrontendURL() string {
	return s.config.FrontendURL
}

func (s *Service) GetGithubAuthURL() string {
	return s.github.AuthCodeURL()
}

func (s *Service) GithubCallback(ctx context.Context, code string) (string, string, error) {
	socialID, err := s.github.GetSocialID(ctx, code)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch user info from GitHub: %w", err)
	}

	user, err := s.database.FindUserBySocialUID(ctx, "github", socialID, true)
	if err != nil {
		return "", "", fmt.Errorf("failed to find or create user: %w", err)
	}

	accessToken, err := s.tokenManager.GenerateAccessToken(user.GetID())
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}
	refreshToken, err := s.tokenManager.GenerateRefreshToken(user.GetID())
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

func (s *Service) RefreshToken(refreshToken string) (string, error) {
	id, err := s.tokenManager.VerifyRefreshToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	newAccessToken, err := s.tokenManager.GenerateAccessToken(id)
	if err != nil {
		return "", fmt.Errorf("failed to generate new access token: %w", err)
	}

	return newAccessToken, nil
}
