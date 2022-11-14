package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golangApi/dao"
	docmodel "golangApi/model/doc"
	requestmodel "golangApi/model/request"
	responsemodel "golangApi/model/response"
	"time"
)

type CompanyTypeInterface interface {
	Create(ctx context.Context, payload requestmodel.CompanyTypeCreate) (err error)

	All(ctx context.Context, query requestmodel.CompanyTypeAll) (result responsemodel.CompanyTypeAll)

	FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.CompanyTypeTable, err error)
}

type companyTypeImp struct {
}

func CompanyType() CompanyTypeInterface {
	return companyTypeImp{}
}

func (ct companyTypeImp) Create(ctx context.Context, payload requestmodel.CompanyTypeCreate) (err error) {

	docBSON := ct.convertToBSON(payload)

	err = dao.CompanyType().Create(ctx, docBSON)

	if err != nil {
		err = errors.New("không thể tạo loại công ty")
	}

	return
}

func (ct companyTypeImp) All(ctx context.Context, query requestmodel.CompanyTypeAll) (result responsemodel.CompanyTypeAll) {

	opts := options.FindOptions{}
	opts.SetLimit(query.Limit)
	opts.SetSkip(query.Page)

	docs := dao.CompanyType().Find(ctx, bson.M{}, &opts)

	for _, doc := range docs {
		result.Data = append(result.Data, ct.convertToJSON(ctx, doc))
	}

	total, err := dao.CompanyType().Count(ctx, bson.M{})
	if err != nil {
		err = errors.New("đã xảy ra lỗi")
	}

	result.Total = total
	result.Limit = query.Limit

	return

}

func (ct companyTypeImp) FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.CompanyTypeTable, err error) {

	doc, err := dao.CompanyType().FindOne(ctx, bson.M{"_id": id})

	if err != nil {
		err = errors.New("không thể tạo công ty")
	}

	result = ct.convertToJSON(ctx, doc)

	return
}

func (ct companyTypeImp) convertToBSON(payload requestmodel.CompanyTypeCreate) (doc docmodel.CompanyType) {
	doc = docmodel.CompanyType{
		ID:        primitive.NewObjectID(),
		Name:      payload.Name,
		Code:      payload.Code,
		CreatedAt: time.Now(),
	}
	return
}

func (ct companyTypeImp) convertToJSON(ctx context.Context, doc docmodel.CompanyType) (res responsemodel.CompanyTypeTable) {
	res = responsemodel.CompanyTypeTable{
		ID:   doc.ID.Hex(),
		Code: doc.Code,
		Name: doc.Name,
	}
	return
}
