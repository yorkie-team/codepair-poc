package main

import (
	"flag"
	"github.com/labstack/gommon/log"
	"github.com/yorkie-team/codepair/backend-go/internal/server"
)

var (
	config     *server.Config
	configPath string
)

func main() {
	codePair, err := server.New(config)
	if err != nil {
		log.Error(err)
	}
	if err := codePair.Start(); err != nil {
		log.Error(err)
	}
}

func init() {
	flag.StringVar(&configPath, "config", "", "Config path")

	if configPath != "" {
		parsed, err := server.NewConfigFromFile(configPath)
		if err != nil {
			log.Error(err)
			return
		}
		config = parsed
	}

	config.EnsureDefaultValue()
}
