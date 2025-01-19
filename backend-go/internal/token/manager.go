package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ErrInvalidToken is returned when the token is invalid or its signature is invalid.
var ErrInvalidToken = errors.New("invalid token signature")

// Manager is responsible for generating and verifying JWT tokens.
type Manager struct {
	config *Config
}

// New creates and returns a new Manager with the given configuration.
func New(config *Config) *Manager {
	return &Manager{config: config}
}

// GenerateAccessToken creates a new access token using the userID as the subject.
func (m *Manager) GenerateAccessToken(userID string) (string, error) {
	return m.generateJWTToken(
		userID,
		m.config.AccessTokenSecret,
		m.config.ParseAccessTokenExpirationTime(),
	)
}

// GenerateRefreshToken creates a new refresh token using the userID as the subject.
func (m *Manager) GenerateRefreshToken(userID string) (string, error) {
	return m.generateJWTToken(
		userID,
		m.config.RefreshTokenSecret,
		m.config.ParseRefreshTokenExpirationTime(),
	)
}

// VerifyAccessToken parses and validates the given access token string,
// returning the userID if valid.
func (m *Manager) VerifyAccessToken(tokenString string) (string, error) {
	return m.parseAndValidateJWTToken(tokenString, m.config.AccessTokenSecret)
}

// VerifyRefreshToken parses and validates the given refresh token string,
// returning the userID if valid.
func (m *Manager) VerifyRefreshToken(tokenString string) (string, error) {
	return m.parseAndValidateJWTToken(tokenString, m.config.RefreshTokenSecret)
}

// generateJWTToken is a helper for generating JWT tokens given the userID,
// the secret, and the token expiration duration.
func (m *Manager) generateJWTToken(userID, secret string, expiry time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// parseAndValidateJWTToken is a helper for parsing and validating a JWT token
// given the raw token string and the secret.
func (m *Manager) parseAndValidateJWTToken(tokenString, secret string) (string, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	// Ensure the token is valid
	if !parsedToken.Valid {
		return "", ErrInvalidToken
	}

	// Extract claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrInvalidToken
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", ErrInvalidToken
	}
	return userID, nil
}
