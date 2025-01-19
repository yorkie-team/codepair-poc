package auth

import "github.com/yorkie-team/codepair/backend-go/internal/domain/auth/github"

const DefaultFrontendURL = "http://localhost:5173"

type Config struct {
	Github      *github.Config `yaml:"github"`
	FrontendURL string         `yaml:"frontendURL"`
}

func (c *Config) EnsureDefaultValue() {
	if c.Github == nil {
		c.Github = &github.Config{}
	}
	c.Github.EnsureDefaultConfig(c.Github)
	if c.FrontendURL == "" {
		c.FrontendURL = DefaultFrontendURL
	}
}
