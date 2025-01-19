package server

import (
	"github.com/labstack/echo/v4"
	"github.com/yorkie-team/codepair/backend-go/internal/database/mongo"
	"github.com/yorkie-team/codepair/backend-go/internal/domain"
	"github.com/yorkie-team/codepair/backend-go/internal/storage"
	"github.com/yorkie-team/codepair/backend-go/internal/storage/minio"
	"github.com/yorkie-team/codepair/backend-go/internal/storage/s3"
	"github.com/yorkie-team/codepair/backend-go/internal/token"
	"github.com/yorkie-team/codepair/backend-go/internal/transport/routes"
	"github.com/yorkie-team/codepair/backend-go/internal/yorkie"
)

type CodePair struct {
	config     *Config
	httpServer *echo.Echo
}

func New(config *Config) (*CodePair, error) {
	db, err := mongo.Dial(config.Mongo)
	if err != nil {
		return nil, err
	}

	var stg storage.Provider
	if config.Storage.Provider == "s3" {
		stg = s3.New(config.Storage.S3)
	} else {
		stg = minio.New(config.Storage.Minio)
	}

	tok := token.New(config.JWT)
	y := yorkie.New(config.Yorkie)

	services := domain.NewServices(config, db, stg, tok, y)
	handlers := domain.NewHandlers(services)
	e := echo.New()
	routes.RegisterRoutes(e, handlers, tok)
	return &CodePair{
		config: config,
	}, nil
}

func (c *CodePair) Start() error {
	return c.httpServer.Start(":" + c.config.Port)
}
