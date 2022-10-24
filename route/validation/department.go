package routevalidation

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	requestmodel "golangApi/model/request"
	"golangApi/util"
)

type Department struct {
}

func (Department) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.DepartmentCreate

		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		if err := payload.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		_, err := primitive.ObjectIDFromHex(payload.Company)

		if err != nil {
			return util.Response400(c, nil, "id không chính xác")
		}

		c.Set("payload", payload)
		return next(c)
	}
}

func (Department) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.DepartmentUpdate
		)

		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		if err := payload.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return util.Response400(c, nil, "id không chính xác")
		}

		_, err = primitive.ObjectIDFromHex(payload.Company)

		if err != nil {
			return util.Response400(c, nil, "id không chính xác")
		}

		c.Set("id", oid)

		c.Set("payload", payload)
		return next(c)
	}
}

func (Department) UpdateActive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.DepartmentUpdateActive
		)

		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		if err := payload.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return util.Response400(c, nil, "id không chính xác")
		}

		c.Set("id", oid)

		c.Set("payload", payload)
		return next(c)
	}
}

func (Department) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query requestmodel.DepartmentAll

		if err := c.Bind(&query); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		c.Set("query", query)
		return next(c)
	}
}

func (Department) Detail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id = c.Param("id")
		)

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return util.Response400(c, nil, "id không chính xác")
		}

		c.Set("id", oid)
		return next(c)
	}
}
