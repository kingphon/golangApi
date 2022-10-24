package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	requestmodel "golangApi/model/request"
	"golangApi/service"
	"golangApi/util"
)

type Drawer struct{}

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
