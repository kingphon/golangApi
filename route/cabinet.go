package route

import (
	"github.com/labstack/echo/v4"
	"golangApi/controller"
	routevalidation "golangApi/route/validation"
)

func cabinet(e *echo.Echo) {
	var (
		g = e.Group("/cabinet")
		c = controller.Cabinet{}
		v = routevalidation.Cabinet{}
	)

	g.GET("", c.All, v.All)

	g.POST("", c.Create, v.Create)

	g.PUT("/:id", c.Update, v.Update)

	g.PATCH("/:id/active", c.UpdateActive, v.UpdateActive)

	g.GET("/:id", c.Detail, v.Detail)
}
