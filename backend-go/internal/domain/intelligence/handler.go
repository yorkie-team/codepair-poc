package intelligence

import (
	"errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(intelligenceService *Service) *Handler {
	return &Handler{
		service: intelligenceService,
	}
}

func (h *Handler) RunFeature(c echo.Context) error {
	return errors.New("not implemented")
}

func (h *Handler) RunFollowUp(c echo.Context) error {
	return errors.New("not implemented")
}
