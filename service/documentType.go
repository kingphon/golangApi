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

type DocumentTypeInterface interface {
	Create(ctx context.Context, payload requestmodel.DocumentTypeCreate) (err error)

	All(ctx context.Context, query requestmodel.DocumentTypeAll) (result responsemodel.DocumentTypeAll)
}

type documentTypeImp struct {
}

func DocumentType() DocumentTypeInterface {
	return documentTypeImp{}
}

func (ct documentTypeImp) Create(ctx context.Context, payload requestmodel.DocumentTypeCreate) (err error) {

	docBSON := ct.convertToBSON(payload)

	err = dao.DocumentType().Create(ctx, docBSON)

	if err != nil {
		err = errors.New("không thể tạo loại công ty")
	}

	return
}

func (ct documentTypeImp) All(ctx context.Context, query requestmodel.DocumentTypeAll) (result responsemodel.DocumentTypeAll) {

	opts := options.FindOptions{}
	opts.SetLimit(query.Limit)
	opts.SetSkip(query.Page)

	docs := dao.DocumentType().Find(ctx, bson.M{}, &opts)

	for _, doc := range docs {
		result.Data = append(result.Data, ct.convertToJSON(ctx, doc))
	}

	total, err := dao.DocumentType().Count(ctx, bson.M{})
	if err != nil {
		err = errors.New("đã xảy ra lỗi")
	}

	result.Total = total
	result.Limit = query.Limit

	return

}

func (ct documentTypeImp) convertToBSON(payload requestmodel.DocumentTypeCreate) (doc docmodel.DocumentType) {
	doc = docmodel.DocumentType{
		ID:        primitive.NewObjectID(),
		Name:      payload.Name,
		Code:      payload.Code,
		CreatedAt: time.Now(),
	}
	return
}

func (ct documentTypeImp) convertToJSON(ctx context.Context, doc docmodel.DocumentType) (res responsemodel.DocumentTypeTable) {
	res = responsemodel.DocumentTypeTable{
		ID:   doc.ID.Hex(),
		Code: doc.Code,
		Name: doc.Name,
	}
	return
}
