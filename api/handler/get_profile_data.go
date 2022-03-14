package handler

import (
	internal "github.com/hawkkiller/wtc-account/internal/database"
	"github.com/hawkkiller/wtc-account/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetProfileData godoc
// @Summary get profile information
// @Description get information from Bearer JWT Access Token(oAuth)
// @Tags Account
// @Accept */*
// @Produce json
// @Param Authorization header string true "H.Authorization: Bearer `token`"
// @Success 200 {object} model.GetProfileInfoResponseOK
// @Success 400 {object} model.GetProfileInfoResponseBR
// @Router /data [get]
func GetProfileData(e echo.Context) error {
	id := e.Get("id").(string)
	userDB := new(model.UserProfile)

	if res := internal.DB.Where("id = ?", id).First(&userDB); res.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, res.Error.Error())
	}

	return e.JSON(http.StatusOK, userDB.ProfileToDB())
}
