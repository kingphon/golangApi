package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	requestmodel "golangApi/model/request"
	"golangApi/service"
	"golangApi/util"
)

type DocumentType struct{}

func (ct DocumentType) All(c echo.Context) error {
	var (
		query = c.Get("query").(requestmodel.DocumentTypeAll)
		ctx   = c.Request().Context()
	)

	data := service.DocumentType().All(ctx, query)

	return util.Response200(c, data, "")
}

func (ct DocumentType) Create(c echo.Context) error {
	var (
		payload = c.Get("payload").(requestmodel.DocumentTypeCreate)
		ctx     = c.Request().Context()
	)

	err := service.DocumentType().Create(ctx, payload)

	if err != nil {
		return util.Response400(c, bson.M{}, err.Error())
	}

	return util.Response200(c, nil, "")
}
