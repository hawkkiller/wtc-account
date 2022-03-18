package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

// CheckJWT Validates JWT token
func CheckJWT(h string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			var id interface{}
			req := c.Request()

			tokenReq := req.Header.Get(h)
			split := strings.Split(tokenReq, " ")

			if len(split) != 2 {
				return echo.NewHTTPError(http.StatusBadRequest, "Maybe token is not Bearer token or it is empty")
			}

			t := split[1]
			claims := jwt.MapClaims{}

			_, err = jwt.ParseWithClaims(t, &claims, func(token *jwt.Token) (interface{}, error) {
				if claims, ok := token.Claims.(*jwt.MapClaims); ok && claims.Valid() == nil {
					return []byte(os.Getenv("SECRET")), nil
				} else {
					return nil, echo.ErrForbidden
				}
			})
			if int64(claims["exp"].(float64)) < time.Now().Add(time.Hour).Unix() && h == "Refresh" {
				return echo.NewHTTPError(http.StatusForbidden, "Token is not refresh")
			}

			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, "Token is expired or invalid")
			}

			if v := claims.Valid(); v != nil {
				return echo.ErrForbidden
			}

			id = claims["jti"]

			if id != nil {
				c.Set("id", id)
				return next(c)
			} else {
				return echo.ErrBadGateway
			}

		}
	}
}
