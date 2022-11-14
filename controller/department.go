package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	requestmodel "golangApi/model/request"
	"golangApi/service"
	"golangApi/util"
)

type Department struct{}

// All godoc
// @tags Department
// @summary Get list
// @security ApiKeyAuth
// @param payload query requestmodel.DepartmentAll true "Payload"
// @router /department [get]
func (d Department) All(c echo.Context) error {
	var (
		query = c.Get("query").(requestmodel.DepartmentAll)
		ctx   = c.Request().Context()
	)

	if query.Limit == 0 {
		query.Limit = 20
	}

	data := service.Department().All(ctx, query)

	return util.Response200(c, data, "")
}

// Create godoc
// @tags Department
// @summary Create
// @security ApiKeyAuth
// @param payload body requestmodel.DepartmentCreate true "Payload"
// @router /department [post]
func (d Department) Create(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.DepartmentCreate)
		ctx     = c.Request().Context()
	)

	err := service.Department().Create(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Update godoc
// @tags Department
// @summary Update
// @security ApiKeyAuth
// @param id   path      string  true  "Department ID"
// @param payload body requestmodel.DepartmentUpdate true "Payload"
// @router /department/{id} [put]
func (d Department) Update(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.DepartmentUpdate)
		id      = c.Get("id").(primitive.ObjectID)
		ctx     = c.Request().Context()
	)

	err := service.Department().Update(ctx, payload, id)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// UpdateActive godoc
// @tags Department
// @summary UpdateActive
// @security ApiKeyAuth
// @param id   path      string  true  "Department ID"
// @param payload body requestmodel.DepartmentUpdateActive true "Payload"
// @router /department/{id}/active [patch]
func (d Department) UpdateActive(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.DepartmentUpdateActive)
		id      = c.Get("id").(primitive.ObjectID)
		ctx     = c.Request().Context()
	)

	err := service.Department().UpdateActive(ctx, payload, id)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Detail godoc
// @tags Department
// @summary Detail
// @security ApiKeyAuth
// @param id   path      string  true  "Department ID"
// @router /department/{id} [get]
func (d Department) Detail(c echo.Context) error {
	var (
		id  = c.Get("id").(primitive.ObjectID)
		ctx = c.Request().Context()
	)

	res, err := service.Department().FindOneWithId(ctx, id)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, res, "")
}
