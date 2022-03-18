package handler

import (
	internal "github.com/hawkkiller/wtc-account/internal/database"
	"github.com/hawkkiller/wtc-account/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

// UpdateProfile godoc
// @Summary update profile
// @Description update profile by passing access token
// @Tags Account
// @Accept json
// @Produce json
// @Param UpdProfileModel body model.UpdateProfileRequest true "update profile model"
// @Param Authorization header string true "H.Authorization: Bearer `token`"
// @Success 200 {object} model.UpdateProfileResponseOK
// @Failure 400 {object} model.UpdateProfileResponseBR
// @Router /update [get]
func UpdateProfile(e echo.Context) error {
	id := e.Get("id")
	updUser := new(model.UpdateProfileRequest)

	err := e.Bind(&updUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userDB := new(model.UserProfile)

	if res := internal.DB.Where("id = ?", id).First(&userDB); res.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, res.Error.Error())
	}

	if updUser.Username != "" {
		userDB.Username = updUser.Username
	}

	if updUser.Sex != "" {
		userDB.Sex = updUser.Sex
	}

	if res := internal.DB.Save(&userDB); res.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, res.Error.Error())
	} else {
		res := model.UpdateProfileResponseOK{
			Message:      "Profile was successfully updated",
			RowsAffected: res.RowsAffected,
		}
		if res.RowsAffected == -1 {
			res.RowsAffected = 0
		}
		return e.JSON(http.StatusOK, res)
	}
}
