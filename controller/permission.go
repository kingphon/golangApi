package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	requestmodel "golangApi/model/request"
	"golangApi/service"
	"golangApi/util"
)

type Permission struct{}

// All godoc
// @tags Permission
// @security ApiKeyAuth
// @summary All
// @param payload query requestmodel.PermissionAll true "Payload"
// @router /permission [get]
func (p Permission) All(c echo.Context) error {
	var (
		query = c.Get("query").(requestmodel.PermissionAll)
		ctx   = c.Request().Context()
	)

	if query.Limit == 0 {
		query.Limit = 20
	}

	data := service.Permission().All(ctx, query)

	return util.Response200(c, data.Data, "")
}

// Create godoc
// @tags Permission
// @summary Create
// @security ApiKeyAuth
// @param payload body requestmodel.PermissionCreate true "Payload"
// @router /permission [post]
func (p Permission) Create(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.PermissionCreate)
		ctx     = c.Request().Context()
	)

	err := service.Permission().Create(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Update godoc
// @tags Permission
// @summary Update
// @security ApiKeyAuth
// @param id   path      string  true  "Permission ID"
// @param payload body requestmodel.PermissionUpdate true "Payload"
// @router /permission/{id} [put]
func (p Permission) Update(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.PermissionUpdate)
		id      = c.Get("id").(primitive.ObjectID)
		ctx     = c.Request().Context()
	)

	err := service.Permission().Update(ctx, payload, id)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}
