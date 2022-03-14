package pkg

import (
	"github.com/golang-jwt/jwt"
	"os"
	"strconv"
	"time"
)

func newJWT(expTime time.Duration, id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expTime).Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        id,
	})
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

// CreateTokens is a function that returns both refresh and access tokens
//
// ACCESS, REFRESH
func CreateTokens(id uint) (string, string) {
	accessToken, _ := newJWT(time.Duration(24)*time.Hour, strconv.Itoa(int(id)))
	refreshToken, _ := newJWT(time.Duration(24)*time.Hour*30, strconv.Itoa(int(id)))

	return accessToken, refreshToken
}
