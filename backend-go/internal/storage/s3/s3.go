package s3

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

type Provider struct {
	client *s3.Client
	config Config
}

// New initializes the Provider with a specified bucket name.
func New(config *Config) *Provider {
	return &Provider{
		config: *config,
	}
}

// CreateUploadURL generates a pre-signed URL for uploading objects.
func (s *Provider) CreateUploadURL(ctx context.Context, key string, contentType string, contentLength int64) (string, error) {
	return "", errors.New("CreateUploadURL not implemented in this version")
}

// CreateDownloadURL generates a pre-signed URL for downloading objects.
func (s *Provider) CreateDownloadURL(ctx context.Context, key string) (string, error) {
	return "", errors.New("CreateDownloadURL not implemented in this version")
}

// Upload uploads an object to S3.
func (s *Provider) Upload(ctx context.Context, key string, content io.Reader, contentType string) error {
	return errors.New("upload not implemented in this version")
}

// Download retrieves an object from S3.
func (s *Provider) Download(ctx context.Context, key string) (io.ReadCloser, error) {
	return nil, errors.New("Download not implemented in this version")
}

// Delete removes an object from S3.
func (s *Provider) Delete(ctx context.Context, key string) error {
	return errors.New("Delete not implemented in this version")
}
