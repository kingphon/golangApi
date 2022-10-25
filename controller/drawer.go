package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	requestmodel "golangApi/model/request"
	"golangApi/service"
	"golangApi/util"
)

type Drawer struct{}

// All godoc
// @tags Drawer
// @summary Get list
// @security ApiKeyAuth
// @param payload query requestmodel.DrawerAll true "Payload"
// @router /drawer [get]
func (co Drawer) All(c echo.Context) error {
	var (
		query = c.Get("query").(requestmodel.DrawerAll)
		ctx   = c.Request().Context()
	)

	if query.Limit == 0 {
		query.Limit = 20
	}

	data := service.Drawer().All(ctx, query)

	return util.Response200(c, data, "")
}

// Create godoc
// @tags Drawer
// @summary Create
// @security ApiKeyAuth
// @param payload body requestmodel.DrawerCreate true "Payload"
// @router /drawer [post]
func (co Drawer) Create(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.DrawerCreate)
		ctx     = c.Request().Context()
	)

	err := service.Drawer().Create(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Update godoc
// @tags Drawer
// @summary Update
// @security ApiKeyAuth
// @param id   path      string  true  "Drawer ID"
// @param payload body requestmodel.DrawerUpdate true "Payload"
// @router /drawer/{id} [put]
func (co Drawer) Update(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.DrawerUpdate)
		ctx     = c.Request().Context()
	)

	err := service.Drawer().Update(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// UpdateActive godoc
// @tags Drawer
// @summary UpdateActive
// @security ApiKeyAuth
// @param id   path      string  true  "Drawer ID"
// @param payload body requestmodel.DrawerUpdateActive true "Payload"
// @router /drawer/{id}/active [patch]
func (co Drawer) UpdateActive(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.DrawerUpdateActive)
		ctx     = c.Request().Context()
	)

	err := service.Drawer().UpdateActive(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Detail godoc
// @tags Drawer
// @summary Detail
// @security ApiKeyAuth
// @param id   path      string  true  "Drawer ID"
// @router /drawer/{id} [get]
func (co Drawer) Detail(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.DrawerDetail)
		ctx     = c.Request().Context()
	)

	res, err := service.Drawer().FindOneWithId(ctx, payload.ID)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, res, "")
}
