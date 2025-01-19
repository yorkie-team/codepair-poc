package domain

import (
	"github.com/yorkie-team/codepair/backend-go/internal/database"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/auth"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/document"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/files"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/intelligence"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/settings"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/user"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/workspace"
	"github.com/yorkie-team/codepair/backend-go/internal/server"
	"github.com/yorkie-team/codepair/backend-go/internal/storage"
	"github.com/yorkie-team/codepair/backend-go/internal/token"
	"github.com/yorkie-team/codepair/backend-go/internal/yorkie"
)

type Services struct {
	Auth         *auth.Service
	User         *user.Service
	Workspace    *workspace.Service
	Document     *document.Service
	File         *files.Service
	Intelligence *intelligence.Service
	Settings     *settings.Service
}

func NewServices(
	cfg *server.Config,
	db database.Database,
	st storage.Provider,
	tk *token.Manager,
	yk *yorkie.Yorkie,
) *Services {
	return &Services{
		Auth:         auth.NewService(cfg.Auth, db, tk),
		User:         user.NewService(db),
		Workspace:    workspace.NewService(db),
		Document:     document.NewService(db, yk),
		File:         files.NewService(st),
		Intelligence: intelligence.NewService(),
		Settings:     settings.NewService(),
	}
}
