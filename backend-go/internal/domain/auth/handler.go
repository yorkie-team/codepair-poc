package auth

import (
	"fmt"
	"github.com/yorkie-team/codepair/backend-go/api/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(authService *Service) *Handler {
	return &Handler{
		service: authService,
	}
}

func (h *Handler) GithubLogin(c echo.Context) error {
	return c.Redirect(http.StatusTemporaryRedirect, h.service.GetGithubAuthURL())
}

func (h *Handler) GithubCallback(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing code parameter")
	}

	access, refresh, err := h.service.GithubCallback(c.Request().Context(), code)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(
		http.StatusFound,
		fmt.Sprintf(
			"%s/auth/callback?access_token=%s&refresh_token=%s",
			h.service.GetFrontendURL(), access, refresh,
		),
	)
}

func (h *Handler) RefreshToken(c echo.Context) error {
	var req dto.RefreshTokenRequestDto
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	accessToken, err := h.service.RefreshToken(req.RefreshToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, dto.RefreshTokenResponseDto{
		NewAccessToken: accessToken,
	})
}
