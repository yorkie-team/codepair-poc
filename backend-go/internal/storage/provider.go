package storage

import (
	"context"
	"io"
)

type Provider interface {
	CreateUploadURL(ctx context.Context, key string, contentType string, contentLength int64) (string, error)
	CreateDownloadURL(ctx context.Context, key string) (string, error)
	Upload(ctx context.Context, key string, content io.Reader, contentType string) error
	Download(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
}
