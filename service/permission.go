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
)

type PermissionInterface interface {
	Create(ctx context.Context, payload requestmodel.PermissionCreate) (err error)

	All(ctx context.Context, query requestmodel.PermissionAll) (result responsemodel.PermissionAll)

	Update(ctx context.Context, payload requestmodel.PermissionUpdate, id primitive.ObjectID) (err error)
}

type permissionImp struct {
}

func Permission() PermissionInterface {
	return permissionImp{}
}

func (c permissionImp) Create(ctx context.Context, payload requestmodel.PermissionCreate) (err error) {

	docBSON := c.convertToBSON(payload)

	err = dao.Permission().Create(ctx, docBSON)

	if err != nil {
		err = errors.New("không thể tạo quyền")
	}

	return
}

func (c permissionImp) Update(ctx context.Context, payload requestmodel.PermissionUpdate, id primitive.ObjectID) (err error) {

	doc, err := dao.Permission().FindOne(ctx, bson.D{{"_id", id}})

	if err != nil {
		return
	}

	docBSON := c.convertToBSONUpdate(payload, doc)

	err = dao.Permission().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật quyền")
	}

	return
}

func (c permissionImp) All(ctx context.Context, query requestmodel.PermissionAll) (result responsemodel.PermissionAll) {

	opts := options.FindOptions{}
	opts.SetLimit(query.Limit)
	opts.SetSkip(query.Page)

	docs := dao.Permission().Find(ctx, bson.M{}, &opts)

	for _, doc := range docs {
		result.Data = append(result.Data, c.convertToJSON(ctx, doc))
	}

	return
}

func (c permissionImp) convertToBSON(payload requestmodel.PermissionCreate) (doc docmodel.Permission) {
	doc = docmodel.Permission{
		ID:   primitive.NewObjectID(),
		Name: payload.Name,
		Code: payload.Code,
	}
	return
}

func (c permissionImp) convertToBSONUpdate(payload requestmodel.PermissionUpdate, doc docmodel.Permission) (result docmodel.Permission) {
	result = docmodel.Permission{
		ID:   doc.ID,
		Name: payload.Name,
		Code: payload.Code,
	}
	return
}

func (c permissionImp) convertToJSON(ctx context.Context, doc docmodel.Permission) (res string) {
	res = doc.Code
	return
}
