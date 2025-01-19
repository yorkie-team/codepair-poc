package yorkie

const (
	DefaultAddress          = "localhost:8080"
	DefaultProjectName      = "default"
	DefaultProjectSecretKey = ""
)

type Config struct {
	address          string `yaml:"address"`
	projectName      string `yaml:"projectName"`
	projectSecretKey string `yaml:"projectSecretKey"`
}

func (c *Config) EnsureDefaultValue() {
	if c.address == "" {
		c.address = DefaultAddress
	}
	if c.projectName == "" {
		c.projectName = DefaultProjectName
	}
	if c.projectSecretKey == "" {
		c.projectSecretKey = DefaultProjectSecretKey
	}
}
