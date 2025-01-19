package s3

const (
	DefaultBucket = "default-storage"
	DefaultRegion = "us-west-2"
	DefaultAccess = "minioadmin"
	DefaultSecret = "minioadmin"
)

type Config struct {
	Bucket    string `yaml:"Bucket"`
	Region    string `yaml:"Region"`
	AccessKey string `yaml:"AccessKey"`
	SecretKey string `yaml:"SecretKey"`
}

func (c *Config) EnsureDefaultValue() {
	if c.Bucket == "" {
		c.Bucket = DefaultBucket
	}
	if c.Region == "" {
		c.Region = DefaultRegion
	}
	if c.AccessKey == "" {
		c.AccessKey = DefaultAccess
	}
	if c.SecretKey == "" {
		c.SecretKey = DefaultSecret
	}
}
