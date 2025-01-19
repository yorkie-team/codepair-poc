package minio

const (
	DefaultBucket   = "default-storage"
	DefaultEndpoint = "localhost:9000"
	DefaultAccess   = "minioadmin"
	DefaultSecret   = "minioadmin"
)

type Config struct {
	Bucket    string `yaml:"Bucket"`
	Endpoint  string `yaml:"Endpoint"`
	AccessKey string `yaml:"AccessKey"`
	SecretKey string `yaml:"SecretKey"`
}

func (c *Config) EnsureDefaultValue() {
	if c.Bucket == "" {
		c.Bucket = DefaultBucket
	}
	if c.Endpoint == "" {
		c.Endpoint = DefaultEndpoint
	}
	if c.AccessKey == "" {
		c.AccessKey = DefaultAccess
	}
	if c.SecretKey == "" {
		c.SecretKey = DefaultSecret
	}
}
