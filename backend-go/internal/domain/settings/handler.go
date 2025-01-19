package settings

import (
	"errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(settingsService *Service) *Handler {
	return &Handler{
		service: settingsService,
	}
}

func (h *Handler) GetSettings(c echo.Context) error {
	return errors.New("not implemented")
}
