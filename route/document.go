package route

import (
	"github.com/labstack/echo/v4"
	"golangApi/controller"
	"golangApi/middleware"
	routevalidation "golangApi/route/validation"
)

func document(e *echo.Echo) {
	var (
		g = e.Group("/document")
		c = controller.Document{}
		v = routevalidation.Document{}
	)

	g.Use(middleware.IsLoggedIn, middleware.Authentication)

	g.GET("", c.All, v.All)

	g.POST("", c.Create, v.Create)

	g.PUT("/:id", c.Update, v.Update)

	g.PATCH("/:id/active", c.UpdateActive, v.UpdateActive)

	g.GET("/:id", c.Detail, v.Detail)
}
