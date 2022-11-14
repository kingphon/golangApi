package route

import (
	"github.com/labstack/echo/v4"
	"golangApi/controller"
	"golangApi/middleware"
	routevalidation "golangApi/route/validation"
)

func documentType(e *echo.Echo) {
	var (
		g = e.Group("/company-type")
		c = controller.DocumentType{}
		v = routevalidation.DocumentType{}
	)

	g.Use(middleware.IsLoggedIn, middleware.Authentication)

	g.GET("", c.All, v.All)

	g.POST("", c.Create, v.Create)
}
