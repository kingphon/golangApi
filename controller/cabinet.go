package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	requestmodel "golangApi/model/request"
	"golangApi/service"
	"golangApi/util"
)

type Cabinet struct{}

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
