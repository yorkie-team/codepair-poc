package intelligence

import (
	"context"
	"errors"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RunFeature(ctx context.Context) error {
	return errors.New("not implemented")
}

func (s *Service) RunFollowUp(ctx context.Context) error {
	return errors.New("not implemented")
}
