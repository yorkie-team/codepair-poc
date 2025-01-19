package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yorkie-team/codepair/backend-go/internal/domain"
	"github.com/yorkie-team/codepair/backend-go/internal/token"
	"github.com/yorkie-team/codepair/backend-go/internal/transport/middleware"
)

func RegisterRoutes(e *echo.Echo, handlers *domain.Handlers, tok *token.Manager) *echo.Echo {
	e.Use(middleware.JWT(tok))
	e.Use(middleware.CORS())

	// Auth routes
	e.POST("/auth/login/github", handlers.Auth.GithubLogin)
	e.GET("/auth/callback/github", handlers.Auth.GithubCallback)
	e.POST("/auth/refresh", handlers.Auth.RefreshToken)

	// User routes
	e.GET("/users", handlers.User.GetProfile)
	e.PUT("/users", handlers.User.UpdateProfile)

	// Workspace routes
	e.POST("/workspaces", handlers.Workspace.Create)
	e.GET("/workspaces", handlers.Workspace.List)
	e.GET("/workspaces/:id", handlers.Workspace.Get)

	// Document routes
	e.POST("/documents", handlers.Document.Create)
	e.PUT("/documents/:id/title", handlers.Document.UpdateTitle)
	e.POST("/documents/:id/share-token", handlers.Document.CreateShareToken)
	e.GET("/documents/share", handlers.Document.FindFromSharingToken)

	// File routes
	e.POST("/files", handlers.File.CreateUploadPresignedURL)
	e.GET("/files/:file_name", handlers.File.CreateDownloadPresignedURL)
	e.POST("/files/export-markdown", handlers.File.ExportMarkdown)

	return e
}
