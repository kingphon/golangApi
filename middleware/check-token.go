package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golangApi/util"
)

var IsLoggedIn = middleware.JWTWithConfig(util.JWTConfig)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := util.GetClaims(c)

		c.Set("claims", claims)
		return next(c)
	}
}

func CheckPermission(c echo.Context, permissions string) (err error) {
	var (
		claims = c.Get("claims").(*util.MyCustomClaims)
	)

	if claims.Data.IsRoot == true {
		return nil
	}
	if isAllow := util.Contains(claims.Data.Department.Permission, permissions); isAllow == true {
		return nil
	}

	return errors.New("bạn không có quyền thực hiện hành động này")
}
