package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	requestmodel "golangApi/model/request"
	"golangApi/service"
	"golangApi/util"
)

type Document struct{}

// All godoc
// @tags Document
// @summary Get list
// @security ApiKeyAuth
// @param payload query requestmodel.DocumentAll true "Payload"
// @router /document [get]
func (co Document) All(c echo.Context) error {
	var (
		query = c.Get("query").(requestmodel.DocumentAll)
		ctx   = c.Request().Context()
	)

	if query.Limit == 0 {
		query.Limit = 20
	}

	data := service.Document().All(ctx, query)

	return util.Response200(c, data, "")
}

// Create godoc
// @tags Document
// @summary Create
// @security ApiKeyAuth
// @param payload body requestmodel.DocumentCreate true "Payload"
// @router /document [post]
func (co Document) Create(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.DocumentCreate)
		ctx     = c.Request().Context()
	)

	err := service.Document().Create(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Update godoc
// @tags Document
// @summary Update
// @security ApiKeyAuth
// @param id   path      string  true  "Document ID"
// @param payload body requestmodel.DocumentUpdate true "Payload"
// @router /document/{id} [put]
func (co Document) Update(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.DocumentUpdate)
		ctx     = c.Request().Context()
	)

	err := service.Document().Update(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// UpdateActive godoc
// @tags Document
// @summary UpdateActive
// @security ApiKeyAuth
// @param id   path      string  true  "Document ID"
// @param payload body requestmodel.DocumentUpdateActive true "Payload"
// @router /document/{id}/active [patch]
func (co Document) UpdateActive(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.DocumentUpdateStatus)
		ctx     = c.Request().Context()
	)

	err := service.Document().UpdateActive(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Detail godoc
// @tags Document
// @summary Detail
// @security ApiKeyAuth
// @param id   path      string  true  "Document ID"
// @router /document/{id} [get]
func (co Document) Detail(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.DocumentDetail)
		ctx     = c.Request().Context()
	)

	res, err := service.Document().FindOneWithId(ctx, payload.ID)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, res, "")
}
