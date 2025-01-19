package storage

import (
	"github.com/yorkie-team/codepair/backend-go/internal/storage/minio"
	"github.com/yorkie-team/codepair/backend-go/internal/storage/s3"
	"log"
	"os"
)

const DefaultProvider = "minio"

type Config struct {
	Provider string        `yaml:"provider"`
	S3       *s3.Config    `yaml:"s3"`
	Minio    *minio.Config `yaml:"minio"`
}

func (c *Config) EnsureDefaultValue() {
	if c.Provider == "" {
		c.Provider = DefaultProvider
	}

	if c.Provider == "s3" && c.S3 == nil {
		c.S3 = &s3.Config{}
		c.S3.EnsureDefaultValue()
	} else if c.Provider == "minio" && c.Minio == nil {
		c.Minio = &minio.Config{}
		c.Minio.EnsureDefaultValue()
	} else {
		log.Println("invalid storage provider")
		os.Exit(1)
	}

}
