package server

import (
	"fmt"
	"github.com/yorkie-team/codepair/backend-go/internal/database/mongo"
	"github.com/yorkie-team/codepair/backend-go/internal/domain/auth"
	"github.com/yorkie-team/codepair/backend-go/internal/storage"
	"github.com/yorkie-team/codepair/backend-go/internal/token"
	"github.com/yorkie-team/codepair/backend-go/internal/yorkie"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const DefaultPort = "8080"

type Config struct {
	Auth    *auth.Config    `yaml:"auth"`
	JWT     *token.Config   `yaml:"jwt"`
	Storage *storage.Config `yaml:"storage"`
	Mongo   *mongo.Config   `yaml:"mongo"`
	Yorkie  *yorkie.Config  `yaml:"yorkie"`
	Port    string          `yaml:"port"`
}

// NewConfigFromFile returns a Config struct for the given conf file.
func NewConfigFromFile(path string) (*Config, error) {
	conf := &Config{}
	bytes, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	if err = yaml.Unmarshal(bytes, conf); err != nil {
		return nil, fmt.Errorf("unmarshal config file: %w", err)
	}

	return conf, nil
}

func (c *Config) EnsureDefaultValue() {
	if c.Auth == nil {
		c.Auth = &auth.Config{}
	}
	c.Auth.EnsureDefaultValue()
	if c.JWT == nil {
		c.JWT = &token.Config{}
	}
	c.JWT.EnsureDefaultConfig()
	if c.Storage == nil {
		c.Storage = &storage.Config{}
	}
	c.Storage.EnsureDefaultValue()
	if c.Mongo == nil {
		c.Mongo = &mongo.Config{}
	}
	c.Mongo.EnsureDefaultValue()
	if c.Yorkie == nil {
		c.Yorkie = &yorkie.Config{}
	}
	c.Yorkie.EnsureDefaultValue()
	if c.Port == "" {
		c.Port = DefaultPort
	}
}
