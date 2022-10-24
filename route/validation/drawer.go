package routevalidation

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	requestmodel "golangApi/model/request"
	"golangApi/util"
)

type Drawer struct {
}

func (Drawer) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.DrawerCreate

		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		if err := payload.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		oid, err := primitive.ObjectIDFromHex(payload.CabinetString)

		if err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		payload.Cabinet = oid

		c.Set("payload", payload)
		return next(c)
	}
}

func (Drawer) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.DrawerUpdate
		)

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

func (Drawer) UpdateActive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.DrawerUpdateActive
		)

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

func (Drawer) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query requestmodel.DrawerAll

		if err := c.Bind(&query); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		c.Set("query", query)
		return next(c)
	}
}

func (Drawer) Detail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.DrawerDetail
		)

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		payload.ID = oid

		c.Set("payload", payload)
		return next(c)
	}
}
