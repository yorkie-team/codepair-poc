package token

import (
	"fmt"
	"os"
	"time"
)

const (
	DefaultAccessTokenSecret          = "access-key"
	DefaultAccessTokenExpirationTime  = "1d"
	DefaultRefreshTokenSecret         = "refresh-key"
	DefaultRefreshTokenExpirationTime = "7d"
)

type Config struct {
	AccessTokenSecret          string `yaml:"accessTokenSecret"`
	AccessTokenExpirationTime  string `yaml:"accessTokenExpirationTime"`
	RefreshTokenSecret         string `yaml:"refreshTokenSecret"`
	RefreshTokenExpirationTime string `yaml:"refreshTokenExpirationTime"`
}

func (c *Config) EnsureDefaultConfig() {
	if c.AccessTokenSecret == "" {
		c.AccessTokenSecret = DefaultAccessTokenSecret
	}
	if c.AccessTokenExpirationTime == "" {
		c.AccessTokenExpirationTime = DefaultAccessTokenExpirationTime
	}
	if c.RefreshTokenSecret == "" {
		c.RefreshTokenSecret = DefaultRefreshTokenSecret
	}
	if c.RefreshTokenExpirationTime == "" {
		c.RefreshTokenExpirationTime = DefaultRefreshTokenExpirationTime
	}
}

func (c *Config) ParseAccessTokenExpirationTime() time.Duration {
	d, err := time.ParseDuration(c.AccessTokenExpirationTime)
	if err != nil {
		fmt.Fprint(os.Stderr, "parse access token expiration time: %w", err)
		os.Exit(1)
	}

	return d
}

func (c *Config) ParseRefreshTokenExpirationTime() time.Duration {
	d, err := time.ParseDuration(c.RefreshTokenExpirationTime)
	if err != nil {
		fmt.Fprint(os.Stderr, "parse refresh token expiration time: %w", err)
		os.Exit(1)
	}

	return d
}
