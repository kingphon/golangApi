package util

import (
	jwt2 "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	responsemodel "golangApi/model/response"
	"time"
)

type MyCustomClaims struct {
	Data responsemodel.StaffTable `json:"data"`
	jwt.RegisteredClaims
}

var mySigningKey = []byte("Keochau")

var JWTConfig = middleware.JWTConfig{
	Claims:     &MyCustomClaims{},
	SigningKey: mySigningKey,
	ContextKey: "staff",
}

func GenerateToken(data responsemodel.StaffTable) (ss string, err error) {
	claims := MyCustomClaims{
		data,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err = token.SignedString(mySigningKey)
	return
}

func GetClaims(c echo.Context) (claims *MyCustomClaims) {
	staff := c.Get("staff").(*jwt2.Token)
	claims = staff.Claims.(*MyCustomClaims)
	return
}
