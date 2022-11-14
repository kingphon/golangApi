package routevalidation

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golangApi/middleware"
	requestmodel "golangApi/model/request"
	"golangApi/util"
)

type Company struct {
}

func (Company) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.CompanyCreate

		if err := middleware.CheckPermission(c, "company_edit"); err != nil {
			return util.Response403(c, nil, err.Error())
		}

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

func (Company) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.CompanyUpdate
		)

		if err := middleware.CheckPermission(c, "company_edit"); err != nil {
			return util.Response403(c, nil, err.Error())
		}

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

func (Company) UpdateActive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.CompanyUpdateActive
		)

		if err := middleware.CheckPermission(c, "company_edit"); err != nil {
			return util.Response403(c, nil, err.Error())
		}

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

func (Company) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			query requestmodel.CompanyAll
		)

		if err := middleware.CheckPermission(c, "company_view"); err != nil {
			return util.Response403(c, nil, err.Error())
		}

		if err := c.Bind(&query); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		c.Set("query", query)
		return next(c)
	}
}

func (Company) Detail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id = c.Param("id")
		)

		if err := middleware.CheckPermission(c, "company_view"); err != nil {
			return util.Response403(c, nil, err.Error())
		}

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return util.Response400(c, nil, "id không chính xác")
		}

		c.Set("id", oid)
		return next(c)
	}
}
