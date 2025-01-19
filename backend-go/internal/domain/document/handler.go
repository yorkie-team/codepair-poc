package document

import (
	"errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(documentService *Service) *Handler {
	return &Handler{
		service: documentService,
	}
}

func (h *Handler) Create(c echo.Context) error {
	return errors.New("not implemented")

}

func (h *Handler) UpdateTitle(c echo.Context) error {
	return errors.New("not implemented")

}

func (h *Handler) CreateShareToken(c echo.Context) error {
	return errors.New("not implemented")

}

func (h *Handler) FindFromSharingToken(c echo.Context) error {
	return errors.New("not implemented")

}
