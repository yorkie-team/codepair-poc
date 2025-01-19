package files

import (
	"errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(fileService *Service) *Handler {
	return &Handler{
		service: fileService,
	}
}

func (h *Handler) CreateUploadPresignedURL(c echo.Context) error {
	return errors.New("not implemented")
}

func (h *Handler) CreateDownloadPresignedURL(c echo.Context) error {
	return errors.New("not implemented")
}

func (h *Handler) ExportMarkdown(c echo.Context) error {
	return errors.New("not implemented")
}
