package yorkie

type Yorkie struct {
	config *Config
}

func New(config *Config) *Yorkie {
	return &Yorkie{
		config: config,
	}
}

func (y *Yorkie) GetDocuments() {

}
