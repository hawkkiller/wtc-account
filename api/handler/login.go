package handler

import (
	"errors"
	internal "github.com/hawkkiller/wtc-account/internal/database"
	"github.com/hawkkiller/wtc-account/internal/model"
	pkg "github.com/hawkkiller/wtc-account/pkg/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

// Login godoc
// @Summary login into an account
// @Description
// @Tags Account
// @Accept json
// @Produce json
// @Param LogProfileModel body model.LogProfileRequest true "user log model"
// @Success 200 {object} model.LogProfileResponseOK
// @Failure 400 {object} model.LogProfileResponseBR
// @Router /login [post]
func Login(e echo.Context) error {
	user := new(model.UserProfile)
	userDB := new(model.UserProfile)
	if err := e.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if res := internal.DB.Where("email = ?", user.Email).First(&userDB); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusBadRequest, gorm.ErrRecordNotFound.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, res.Error.Error())

	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err == nil {
			accessToken, refreshToken := pkg.CreateTokens(userDB.ID)
			res := model.LogProfileResponseOK{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			}

			return e.JSON(http.StatusOK, res)
		} else {
			res := model.LogProfileResponseBR{Message: "Password is not correct."}
			return e.JSON(http.StatusBadRequest, res)
		}

	}
}
