package routevalidation

import (
	"github.com/labstack/echo/v4"
	requestmodel "golangApi/model/request"
	"golangApi/util"
)

type DocumentType struct {
}

func (DocumentType) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.DocumentTypeCreate

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

func (DocumentType) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query requestmodel.DocumentTypeAll

		if err := c.Bind(&query); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		if err := query.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("query", query)
		return next(c)
	}
}
