package config

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yumekiti/cocoroiki-bff/domain"
)

type JwtCustomClaims struct {
	domain.User
	jwt.StandardClaims
}

func MD5(text string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

// {id: string, pw: string}
type Param struct {
	ID       string `json:"id"`
	Password string `json:"pw"`
}

func Login(c echo.Context) error {
	// Bind
	param := new(Param)
	if err := c.Bind(param); err != nil {
		return err
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		User: domain.User{
			ID:       param.ID,
			Password: param.Password,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 2147483647,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func JWTConfig() *middleware.JWTConfig {
	return &middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
}
