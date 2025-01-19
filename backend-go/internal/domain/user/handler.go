package user

import (
	"errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(userService *Service) *Handler {
	return &Handler{
		service: userService,
	}
}

func (h *Handler) GetProfile(c echo.Context) error {
	return errors.New("not implemented")
}

func (h *Handler) UpdateProfile(c echo.Context) error {
	return errors.New("not implemented")
}
