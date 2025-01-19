package workspace

import (
	"errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	workspaceService *Service
}

func NewHandler(workspaceService *Service) *Handler {
	return &Handler{
		workspaceService: workspaceService,
	}
}

func (h *Handler) Create(c echo.Context) error {
	return errors.New("not implemented")
}

func (h *Handler) List(c echo.Context) error {
	return errors.New("not implemented")
}

func (h *Handler) Get(c echo.Context) error {
	return errors.New("not implemented")
}

func (h *Handler) CreateInvitationToken(c echo.Context) error {
	return errors.New("not implemented")
}
