package database

import (
	"context"
)

// Database defines all DB methods used throughout the services.
// Methods are grouped by domain (e.g., Users, Workspaces, Documents, etc.).
type Database interface {
	// FindUserByID returns a user by its ID, or nil if not found.
	FindUserByID(ctx context.Context, userID string) (*User, error)

	// FindUserBySocialUID checks if a user exists with given provider & uid.
	FindUserBySocialUID(ctx context.Context, provider, socialID string, createIfNotExist bool) (*User, error)

	// Ping is a fallback method if needed to check connectivity.
	Ping(ctx context.Context) error
}
