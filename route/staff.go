package route

import (
	"github.com/labstack/echo/v4"
	"golangApi/controller"
	routevalidation "golangApi/route/validation"
)

func staff(e *echo.Echo) {
	var (
		g = e.Group("/staff")
		c = controller.Staff{}
		v = routevalidation.Staff{}
	)

	g.GET("", c.All, v.All)

	g.POST("", c.Create, v.Create)

	g.PUT("/:id", c.Update, v.Update)

	g.PATCH("/:id/active", c.UpdateActive, v.UpdateActive)

	g.PATCH("/:id/password", c.UpdatePassword, v.UpdatePassword)

	e.POST("/login", c.Login, v.Login)

	g.GET("/:id", c.Detail, v.Detail)
}
