package routevalidation

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golangApi/middleware"
	requestmodel "golangApi/model/request"
	"golangApi/util"
)

type Cabinet struct {
}

func (Cabinet) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.CabinetCreate

		if err := middleware.CheckPermission(c, "cabinet_edit"); err != nil {
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

func (Cabinet) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.CabinetUpdate
		)

		if err := middleware.CheckPermission(c, "cabinet_edit"); err != nil {
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
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		payload.ID = oid

		c.Set("payload", payload)
		return next(c)
	}
}

func (Cabinet) UpdateActive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.CabinetUpdateActive
		)

		if err := middleware.CheckPermission(c, "cabinet_edit"); err != nil {
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
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		payload.ID = oid

		c.Set("payload", payload)
		return next(c)
	}
}

func (Cabinet) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query requestmodel.CabinetAll

		if err := middleware.CheckPermission(c, "cabinet_view"); err != nil {
			return util.Response403(c, nil, err.Error())
		}

		if err := c.Bind(&query); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		c.Set("query", query)
		return next(c)
	}
}

func (Cabinet) Detail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.CabinetDetail
		)

		if err := middleware.CheckPermission(c, "cabinet_view"); err != nil {
			return util.Response403(c, nil, err.Error())
		}

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		payload.ID = oid

		c.Set("payload", payload)
		return next(c)
	}
}
