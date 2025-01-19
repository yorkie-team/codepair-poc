package mongo

import (
	"fmt"
	"os"
	"time"
)

const (
	// DefaultConnectionTimeout is the default connection timeout.
	DefaultConnectionTimeout = "10s"
	// DefaultPingTimeout is the default ping timeout.
	DefaultPingTimeout = "10s"
	// DefaultDatabaseName is the default database name.
	DefaultDatabaseName = "codepair"
	// DefaultConnectionURI is the default connection URI.
	DefaultConnectionURI = "mongodb://localhost:27017/codepair"
)

type Config struct {
	ConnectionTimeout string `yaml:"ConnectionTimeout"`
	ConnectionURI     string `yaml:"ConnectionURI"`
	PingTimeout       string `yaml:"PingTimeout"`
	DatabaseName      string `yaml:"DatabaseName"`
}

func (c *Config) ParseConnectionTimeout() time.Duration {
	result, err := time.ParseDuration(c.ConnectionTimeout)
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse connection timeout: %w", err)
		os.Exit(1)
	}

	return result
}

// ParsePingTimeout returns ping timeout duration.
func (c *Config) ParsePingTimeout() time.Duration {
	result, err := time.ParseDuration(c.PingTimeout)
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse ping timeout: %w", err)
		os.Exit(1)
	}

	return result
}

// EnsureDefaultValue ensures the default value.
func (c *Config) EnsureDefaultValue() {
	if c.ConnectionTimeout == "" {
		c.ConnectionTimeout = DefaultConnectionTimeout
	}
	if c.ConnectionURI == "" {
		c.ConnectionURI = DefaultConnectionURI
	}
	if c.PingTimeout == "" {
		c.PingTimeout = DefaultPingTimeout
	}
	if c.DatabaseName == "" {
		c.DatabaseName = DefaultDatabaseName
	}
}
