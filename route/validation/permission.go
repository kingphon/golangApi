package routevalidation

import (
	"github.com/labstack/echo/v4"
	requestmodel "golangApi/model/request"
	"golangApi/util"
)

type Permission struct {
}

func (Permission) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.PermissionCreate

		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		if err := payload.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("payload", payload)

		return next(c)
	}
}

func (Permission) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload requestmodel.PermissionUpdate
		)

		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		if err := payload.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("payload", payload)
		return next(c)
	}
}
func (Permission) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			query requestmodel.PermissionAll
		)

		if err := c.Bind(&query); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		c.Set("query", query)
		return next(c)
	}
}
