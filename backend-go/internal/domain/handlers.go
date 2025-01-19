package domain

import (
	"github.com/yorkie-team/codepair/backend-go/internal/domain/auth"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/document"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/files"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/intelligence"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/settings"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/user"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/workspace"
)

type Handlers struct {
	Auth         *auth.Handler
	User         *user.Handler
	Workspace    *workspace.Handler
	Document     *document.Handler
	File         *files.Handler
	Intelligence *intelligence.Handler
	Settings     *settings.Handler
}

func NewHandlers(services *Services) *Handlers {
	return &Handlers{
		Auth:         auth.NewHandler(services.Auth),
		User:         user.NewHandler(services.User),
		Workspace:    workspace.NewHandler(services.Workspace),
		Document:     document.NewHandler(services.Document),
		File:         files.NewHandler(services.File),
		Intelligence: intelligence.NewHandler(services.Intelligence),
		Settings:     settings.NewHandler(services.Settings),
	}
}
