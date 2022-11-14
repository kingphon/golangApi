package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	requestmodel "golangApi/model/request"
	"golangApi/service"
	"golangApi/util"
)

type Company struct{}

// All godoc
// @tags Company
// @security ApiKeyAuth
// @summary Get list
// @param payload query requestmodel.CompanyAll true "Payload"
// @router /company [get]
func (co Company) All(c echo.Context) error {
	var (
		query = c.Get("query").(requestmodel.CompanyAll)
		ctx   = c.Request().Context()
	)

	if query.Limit == 0 {
		query.Limit = 20
	}

	data := service.Company().All(ctx, query)

	return util.Response200(c, data, "")
}

// Create godoc
// @tags Company
// @summary Create
// @security ApiKeyAuth
// @param payload body requestmodel.CompanyCreate true "Payload"
// @router /company [post]
func (co Company) Create(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.CompanyCreate)
		ctx     = c.Request().Context()
	)

	err := service.Company().Create(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Update godoc
// @tags Company
// @summary Update
// @security ApiKeyAuth
// @param id   path      string  true  "Company ID"
// @param payload body requestmodel.CompanyUpdate true "Payload"
// @router /company/{id} [put]
func (co Company) Update(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.CompanyUpdate)
		id      = c.Get("id").(primitive.ObjectID)
		ctx     = c.Request().Context()
	)

	err := service.Company().Update(ctx, payload, id)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// UpdateActive godoc
// @tags Company
// @summary UpdateActive
// @security ApiKeyAuth
// @param id   path      string  true  "Company ID"
// @param payload body requestmodel.CompanyUpdateActive true "Payload"
// @router /company/{id}/active [patch]
func (co Company) UpdateActive(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.CompanyUpdateActive)
		id      = c.Get("id").(primitive.ObjectID)
		ctx     = c.Request().Context()
	)

	err := service.Company().UpdateActive(ctx, payload, id)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Detail godoc
// @tags Company
// @summary Detail
// @security ApiKeyAuth
// @param id   path      string  true  "Company ID"
// @router /company/{id} [get]
func (co Company) Detail(c echo.Context) error {
	var (
		id  = c.Get("id").(primitive.ObjectID)
		ctx = c.Request().Context()
	)

	res, err := service.Company().FindOneWithId(ctx, id)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, res, "")
}
