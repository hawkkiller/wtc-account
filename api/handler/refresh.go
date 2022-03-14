package handler

import (
	"github.com/hawkkiller/wtc-account/internal/model"
	pkg "github.com/hawkkiller/wtc-account/pkg/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Refresh godoc
// @Summary create new access token using refresh token
// @Description pass refresh token to header `Refresh`
// @Param Refresh header string true "refresh token"
// @Tags Account
// @Accept json
// @Produce json
// @Success 200 {object} model.ReAuthResponseOK
// @Failure 400 {object} model.ReAuthResponseBR
// @Failure 403 {object} model.ReAuthResponseBR
// @Router /refresh [get]
func Refresh(e echo.Context) error {
	id, err := strconv.Atoi(e.Get("id").(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	accessT, refreshT := pkg.CreateTokens(uint(id))

	return e.JSON(http.StatusOK, model.ReAuthResponseOK{
		AccessToken:  accessT,
		RefreshToken: refreshT,
	})
}
