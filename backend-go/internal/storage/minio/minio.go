package minio

import (
	"context"
	"errors"
	"io"
)

type Provider struct {
	config *Config
}

func New(config *Config) *Provider {
	return &Provider{
		config: config,
	}
}

func (s *Provider) CreateUploadURL(ctx context.Context, key string, contentType string, contentLength int64) (string, error) {
	return "", errors.New("CreateUploadURL not implemented in this version")
}

func (s *Provider) CreateDownloadURL(ctx context.Context, key string) (string, error) {
	return "", errors.New("CreateDownloadURL not implemented in this version")
}

func (s *Provider) Upload(ctx context.Context, key string, content io.Reader, contentType string) error {
	return errors.New("Upload not implemented in this version")
}

func (s *Provider) Download(ctx context.Context, key string) (io.ReadCloser, error) {
	return nil, errors.New("Download not implemented in this version")
}

func (s *Provider) Delete(ctx context.Context, key string) error {
	return errors.New("Delete not implemented in this version")
}
