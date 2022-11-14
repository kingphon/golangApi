package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	requestmodel "golangApi/model/request"
	"golangApi/service"
	"golangApi/util"
)

type CompanyType struct{}

func (ct CompanyType) All(c echo.Context) error {
	var (
		query = c.Get("query").(requestmodel.CompanyTypeAll)
		ctx   = c.Request().Context()
	)

	data := service.CompanyType().All(ctx, query)

	return util.Response200(c, data, "")
}

func (ct CompanyType) Create(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.CompanyTypeCreate)
		ctx     = c.Request().Context()
	)

	err := service.CompanyType().Create(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}
