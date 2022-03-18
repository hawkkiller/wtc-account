package api

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/hawkkiller/wtc-account/transport/httpApi/handler"
	customMiddleware "github.com/hawkkiller/wtc-account/transport/httpApi/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
)

type AccountServerHTTP struct {
	*echo.Echo
	Port string
}
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

// NewServerHTTP Applies middlewares and sets handlers
func NewServerHTTP() (s *AccountServerHTTP) {
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	return &AccountServerHTTP{e, port}
}

func (s *AccountServerHTTP) StartServerHTTP() error {

	port := fmt.Sprintf(":%s", s.Port)
	fmt.Println("Started HTTP server on port ", port)

	return s.Start(port)
}
