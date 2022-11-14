package routevalidation

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golangApi/middleware"
	requestmodel "golangApi/model/request"
	"golangApi/util"
)

type Document struct {
}

func (Document) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.DocumentCreate

		if err := middleware.CheckPermission(c, "document_edit"); err != nil {
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

func (Document) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.DocumentUpdate
		)

		if err := middleware.CheckPermission(c, "document_edit"); err != nil {
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

func (Document) UpdateActive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.DocumentUpdateStatus
		)

		if err := middleware.CheckPermission(c, "document_edit"); err != nil {
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

func (Document) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query requestmodel.DocumentAll

		if err := c.Bind(&query); err != nil {
			return util.Response400(c, nil, "đã xảy ra lỗi")
		}

		if err := middleware.CheckPermission(c, "document_view"); err != nil {
			return util.Response403(c, nil, err.Error())
		}

		c.Set("query", query)
		return next(c)
	}
}

func (Document) Detail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload requestmodel.DocumentDetail
		)

		if err := middleware.CheckPermission(c, "document_view"); err != nil {
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
