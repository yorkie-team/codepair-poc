package files

import (
	"errors"
	"github.com/yorkie-team/codepair/backend-go/internal/storage"
)

type Service struct {
	storage storage.Provider
}

func NewService(storage storage.Provider) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateUploadPresignedURL() (string, error) {
	return "", errors.New("not implemented")
}

func (s *Service) CreateDownloadPresignedURL(fileKey string) (string, error) {

	return "", errors.New("not implemented")
}

func (s *Service) ExportMarkdown() error {
	return errors.New("not implemented")
}
