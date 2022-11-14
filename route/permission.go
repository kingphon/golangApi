package route

import (
	"github.com/labstack/echo/v4"
	"golangApi/controller"
	"golangApi/middleware"
	routevalidation "golangApi/route/validation"
)

func permission(e *echo.Echo) {
	var (
		g = e.Group("/permission")
		c = controller.Permission{}
		v = routevalidation.Permission{}
	)

	g.Use(middleware.IsLoggedIn, middleware.Authentication)

	g.GET("", c.All, v.All)

	g.POST("", c.Create, v.Create)

	g.PUT("/:id", c.Update, v.Update)
}
