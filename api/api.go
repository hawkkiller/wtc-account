package api

import (
	"github.com/hawkkiller/wtc-account/api/handler"
	customMiddleware "github.com/hawkkiller/wtc-account/api/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type AccountService struct {
	*echo.Echo
}

// CreateApi Applies middlewares and sets handlers
func CreateApi() (s *AccountService) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api/v1/account-service")

	api.GET("/swagger/*", echoSwagger.WrapHandler)

	api.GET("/data", handler.GetProfileData, customMiddleware.CheckJWT("Authorization"))

	api.PUT("/update", handler.UpdateProfile, customMiddleware.CheckJWT("Authorization"))

	api.GET("/refresh", handler.Refresh, customMiddleware.CheckJWT("Refresh"))

	api.POST("/login", handler.Login)

	api.POST("/register", handler.Register)

	return &AccountService{e}
}
