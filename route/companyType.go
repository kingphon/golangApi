package route

import (
	"github.com/labstack/echo/v4"
	"golangApi/controller"
	"golangApi/middleware"
	routevalidation "golangApi/route/validation"
)

func companyType(e *echo.Echo) {
	var (
		g = e.Group("/company-type")
		c = controller.CompanyType{}
		v = routevalidation.CompanyType{}
	)

	g.Use(middleware.IsLoggedIn, middleware.Authentication)

	g.GET("", c.All, v.All)

	g.POST("", c.Create, v.Create)
}
