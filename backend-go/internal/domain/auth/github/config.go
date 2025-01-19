package github

import (
	"fmt"
	"os"
)

const (
	DefaultCallbackURL      = "http://localhost:3000/auth/login/github"
	DefaultAuthorizationURL = "https://github.com/login/oauth/authorize"
	DefaultTokenURL         = "https://github.com/login/oauth/access_token"
	DefaultUserProfileURL   = "https://api.github.com/user"
)

type Config struct {
	ClientID         string `yaml:"clientID"`
	ClientSecret     string `yaml:"clientSecret"`
	CallbackURL      string `yaml:"callbackURL"`
	AuthorizationURL string `yaml:"authorizationURL"`
	TokenURL         string `yaml:"tokenURL"`
	UserProfileURL   string `yaml:"userProfileURL"`
}

func (c *Config) EnsureDefaultConfig(config *Config) {
	if config.ClientID == "" {
		fmt.Println("Error: GitHub Client ID is not set.")
		os.Exit(1)
	}
	if config.ClientSecret == "" {
		fmt.Println("Error: GitHub Client Secret is not set.")
		os.Exit(1)
	}
	if config.CallbackURL == "" {
		config.CallbackURL = DefaultCallbackURL
	}
	if config.AuthorizationURL == "" {
		config.AuthorizationURL = DefaultAuthorizationURL
	}
	if config.TokenURL == "" {
		config.TokenURL = DefaultTokenURL
	}
	if config.UserProfileURL == "" {
		config.UserProfileURL = DefaultUserProfileURL
	}
}
