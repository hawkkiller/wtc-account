package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/hawkkiller/wtc-account/api/handler"
	customMiddleware "github.com/hawkkiller/wtc-account/api/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// SetupApi Applies middlewares and sets handlers
func SetupApi() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &CustomValidator{validator: validator.New()}

	api := e.Group("/api/v1/account-service")

	api.GET("/swagger/*", echoSwagger.WrapHandler)

	api.GET("/data", handler.GetProfileData, customMiddleware.CheckJWT("Authorization"))

	api.PUT("/update", handler.UpdateProfile, customMiddleware.CheckJWT("Authorization"))

	api.GET("/refresh", handler.Refresh, customMiddleware.CheckJWT("Refresh"))

	api.POST("/login", handler.Login)

	api.POST("/register", handler.Register)

	return e
}
