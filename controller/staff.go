package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	requestmodel "golangApi/model/request"
	"golangApi/service"
	"golangApi/util"
)

type Staff struct{}

// All godoc
// @tags Staff
// @summary Get list
// @security ApiKeyAuth
// @param payload query requestmodel.StaffAll true "Payload"
// @router /staff [get]
func (s Staff) All(c echo.Context) error {
	var (
		query = c.Get("query").(requestmodel.StaffAll)
		ctx   = c.Request().Context()
	)

	if query.Limit == 0 {
		query.Limit = 20
	}

	data := service.Staff().All(ctx, query)

	return util.Response200(c, data, "")
}

// Create godoc
// @tags Staff
// @summary Create
// @security ApiKeyAuth
// @param payload body requestmodel.StaffCreate true "Payload"
// @router /staff [post]
func (s Staff) Create(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.StaffCreate)
		ctx     = c.Request().Context()
	)

	err := service.Staff().Create(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Update godoc
// @tags Staff
// @summary Update
// @security ApiKeyAuth
// @param id   path      string  true  "Staff ID"
// @param payload body requestmodel.StaffUpdate true "Payload"
// @router /staff/{id} [put]
func (s Staff) Update(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.StaffUpdate)
		id      = c.Get("id").(primitive.ObjectID)
		ctx     = c.Request().Context()
	)

	err := service.Staff().Update(ctx, payload, id)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// UpdateActive godoc
// @tags Staff
// @summary UpdateActive
// @security ApiKeyAuth
// @param id   path      string  true  "Staff ID"
// @param payload body requestmodel.StaffUpdateActive true "Payload"
// @router /staff/{id}/active [patch]
func (s Staff) UpdateActive(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.StaffUpdateActive)
		id      = c.Get("id").(primitive.ObjectID)
		ctx     = c.Request().Context()
	)

	err := service.Staff().UpdateActive(ctx, payload, id)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Detail godoc
// @tags Staff
// @summary Detail
// @security ApiKeyAuth
// @param id   path      string  true  "Staff ID"
// @router /staff/{id} [get]
func (s Staff) Detail(c echo.Context) error {
	var (
		id  = c.Get("id").(primitive.ObjectID)
		ctx = c.Request().Context()
	)

	res, err := service.Staff().FindOneWithId(ctx, id)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, res, "")
}

// UpdatePassword godoc
// @tags Staff
// @summary UpdatePassword
// @security ApiKeyAuth
// @param id   path      string  true  "Staff ID"
// @param payload body requestmodel.StaffUpdatePassword true "Payload"
// @router /staff/{id}/password [patch]
func (s Staff) UpdatePassword(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.StaffUpdatePassword)
		id      = c.Get("id").(primitive.ObjectID)
		ctx     = c.Request().Context()
	)

	err := service.Staff().UpdatePassword(ctx, payload, id)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}

// Login godoc
// @tags Staff
// @summary Login
// @param payload body requestmodel.StaffLogin true "Payload"
// @router /login [post]
func (s Staff) Login(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.StaffLogin)
		ctx     = c.Request().Context()
	)

	token, err := service.Staff().Login(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, token, "")
}
