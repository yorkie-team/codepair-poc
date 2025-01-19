package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/yorkie-team/codepair/backend-go/internal/token"
	"net/http"
	"strings"
)

type AuthUserKey struct{}

var PublicPaths = []string{
	"/auth/login/github",
	"/auth/callback/github",
	"/auth/refresh",
	"/check/name-conflict",
	"/check/yorkie",
	"/documents/share",
}

func JWT(manager *token.Manager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Skip auth for public endpoints
			if isPublicPath(c.Path()) {
				return next(c)
			}

			// Get Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header required")
			}

			// Extract token string. Example header: "Bearer <token>"
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return echo.NewHTTPError(http.StatusUnauthorized, "Bearer token is required")
			}

			// Verify token and extract user info
			userID, err := manager.VerifyAccessToken(tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			// Add user info to context
			c.Set("user", userID)

			// Call the next handler
			return next(c)
		}
	}
}
func isPublicPath(path string) bool {
	for _, p := range PublicPaths {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}
