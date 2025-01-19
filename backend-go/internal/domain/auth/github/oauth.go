package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
)

type OAuth struct {
	github *oauth2.Config
	config *Config
}

// UserInfo is the minimal user data we need from GitHub.
type UserInfo struct {
	ID int `json:"id"`
}

func New(config *Config) *OAuth {
	return &OAuth{
		github: &oauth2.Config{
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			RedirectURL:  config.CallbackURL,
			Endpoint: oauth2.Endpoint{
				AuthURL:  config.AuthorizationURL,
				TokenURL: config.TokenURL,
			},
			Scopes: []string{"user:public_profile"},
		},
		config: config,
	}
}

func (o *OAuth) AuthCodeURL() string {
	return o.github.AuthCodeURL("state")
}

// GetSocialID fetches the user's GitHub profile using the given token.
func (o *OAuth) GetSocialID(ctx context.Context, code string) (string, error) {
	token, err := o.github.Exchange(ctx, code)
	if err != nil {
		return "", fmt.Errorf("failed to exchange code for token: %w", err)
	}

	client := o.github.Client(ctx, token)

	resp, err := client.Get(o.config.UserProfileURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch user profile: %w", err)
	}
	defer func(Body io.ReadCloser) {
		if closeErr := Body.Close(); closeErr != nil {
			log.Printf("error closing response body: %v", closeErr)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	var user UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", fmt.Errorf("failed to decode user profile: %w", err)
	}

	if user.ID == 0 {
		return "", errors.New("user ID is missing in the response")
	}

	return fmt.Sprint(user.ID), nil
}
