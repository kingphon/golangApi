package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	requestmodel "golangApi/model/request"
	"golangApi/service"
	"golangApi/util"
)

type Cabinet struct{}

// All godoc
// @tags Cabinet
// @summary Get list
// @security ApiKeyAuth
// @param payload query requestmodel.CabinetAll true "Payload"
// @router /cabinet [get]
func (co Cabinet) All(c echo.Context) error {
	var (
		query = c.Get("query").(requestmodel.CabinetAll)
		ctx   = c.Request().Context()
	)

	if query.Limit == 0 {
		query.Limit = 20
	}

	data := service.Cabinet().All(ctx, query)

	return util.Response200(c, data, "")
}

// Create godoc
// @tags Cabinet
// @summary Create
// @security ApiKeyAuth
// @param payload body requestmodel.CabinetCreate true "Payload"
// @router /cabinet [post]
func (co Cabinet) Create(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.CabinetCreate)
		ctx     = c.Request().Context()
	)

	err := service.Cabinet().Create(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Update godoc
// @tags Cabinet
// @summary Update
// @security ApiKeyAuth
// @param id   path      string  true  "Cabinet ID"
// @param payload body requestmodel.CabinetUpdate true "Payload"
// @router /cabinet/{id} [put]
func (co Cabinet) Update(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.CabinetUpdate)
		ctx     = c.Request().Context()
	)

	err := service.Cabinet().Update(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// UpdateActive godoc
// @tags Cabinet
// @summary UpdateActive
// @security ApiKeyAuth
// @param id   path      string  true  "Cabinet ID"
// @param payload body requestmodel.CabinetUpdateActive true "Payload"
// @router /cabinet/{id}/active [patch]
func (co Cabinet) UpdateActive(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.CabinetUpdateActive)
		ctx     = c.Request().Context()
	)

	err := service.Cabinet().UpdateActive(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Detail godoc
// @tags Cabinet
// @summary Detail
// @security ApiKeyAuth
// @param id   path      string  true  "Cabinet ID"
// @router /cabinet/{id} [get]
func (co Cabinet) Detail(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.CabinetDetail)
		ctx     = c.Request().Context()
	)

	res, err := service.Cabinet().FindOneWithId(ctx, payload.ID)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, res, "")
}
