package routevalidation

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	requestmodel "golangApi/model/request"
	"golangApi/util"
)

type Staff struct {
}

func (Staff) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.StaffCreate

		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		if err := payload.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		_, err := primitive.ObjectIDFromHex(payload.Department)

		if err != nil {
			return util.Response400(c, nil, "id không hợp lệ")
		}

		c.Set("payload", payload)
		return next(c)
	}
}

func (Staff) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.StaffUpdate
		)

		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		if err := payload.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return util.Response400(c, nil, "id không hợp lệ")
		}

		_, err = primitive.ObjectIDFromHex(payload.Department)

		if err != nil {
			return util.Response400(c, nil, "id không hợp lệ")
		}

		c.Set("id", oid)

		c.Set("payload", payload)
		return next(c)
	}
}

func (Staff) UpdateActive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.StaffUpdateActive
		)

		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		if err := payload.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return util.Response400(c, nil, "id không hợp lệ")
		}

		c.Set("id", oid)

		c.Set("payload", payload)
		return next(c)
	}
}

func (Staff) UpdatePassword(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.StaffUpdatePassword
		)

		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		if err := payload.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return util.Response400(c, nil, "id không hợp lệ")
		}

		c.Set("id", oid)

		c.Set("payload", payload)
		return next(c)
	}
}

func (Staff) Login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload requestmodel.StaffLogin
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

func (Staff) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query requestmodel.StaffAll

		if err := c.Bind(&query); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		c.Set("query", query)
		return next(c)
	}
}

func (Staff) Detail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id = c.Param("id")
		)

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return util.Response400(c, nil, "id không hợp lệ")
		}

		c.Set("id", oid)
		return next(c)
	}
}
