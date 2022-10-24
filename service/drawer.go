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

type DrawerInterface interface {
	Create(ctx context.Context, payload requestmodel.DrawerCreate) (err error)

	All(ctx context.Context, query requestmodel.DrawerAll) (result responsemodel.DrawerAll)

	UpdateActive(ctx context.Context, payload requestmodel.DrawerUpdateActive) (err error)

	Update(ctx context.Context, payload requestmodel.DrawerUpdate) (err error)

	FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.DrawerTable, err error)
}

type drawerImp struct {
}

func Drawer() DrawerInterface {
	return drawerImp{}
}

func (c drawerImp) Create(ctx context.Context, payload requestmodel.DrawerCreate) (err error) {

	count, err := dao.Cabinet().Count(ctx, bson.M{"_id": payload.Cabinet})

	if count == 0 {
		err = errors.New("kệ không tồn tại")
		return
	}

	docBSON := c.convertToBSON(payload)

	err = dao.Drawer().Create(ctx, docBSON)

	if err != nil {
		err = errors.New("không thể tạo ngăn")
	}

	return
}

func (c drawerImp) UpdateActive(ctx context.Context, payload requestmodel.DrawerUpdateActive) (err error) {

	doc, err := dao.Drawer().FindOne(ctx, bson.M{"_id": payload.ID})

	if err != nil {
		return
	}

	docBSON := c.convertToBSONUpdateActive(payload, doc)

	err = dao.Drawer().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật ngăn")
	}

	return
}

func (c drawerImp) Update(ctx context.Context, payload requestmodel.DrawerUpdate) (err error) {

	doc, err := dao.Drawer().FindOne(ctx, bson.D{{"_id", payload.ID}})

	if err != nil {
		return
	}

	docBSON := c.convertToBSONUpdate(payload, doc)

	err = dao.Drawer().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật ngăn")
	}

	return
}

func (c drawerImp) All(ctx context.Context, query requestmodel.DrawerAll) (result responsemodel.DrawerAll) {

	opts := options.FindOptions{}
	opts.SetLimit(query.Limit)
	opts.SetSkip(query.Page)

	docs := dao.Drawer().Find(ctx, bson.M{}, &opts)

	for _, doc := range docs {
		result.Data = append(result.Data, c.convertToJSON(ctx, doc))
	}

	total, err := dao.Drawer().Count(ctx, bson.M{})
	if err != nil {
		err = errors.New("đã xảy ra lỗi")
	}

	result.Total = total
	result.Limit = query.Limit

	return

}

func (c drawerImp) FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.DrawerTable, err error) {

	doc, err := dao.Drawer().FindOne(ctx, bson.M{"_id": id})

	if err != nil {
		err = errors.New("không thể tạo ngăn")
	}

	result = c.convertToJSON(ctx, doc)

	return
}

func (c drawerImp) convertToBSON(payload requestmodel.DrawerCreate) (doc docmodel.Drawer) {
	doc = docmodel.Drawer{
		ID:        primitive.NewObjectID(),
		Name:      payload.Name,
		Cabinet:   payload.Cabinet,
		Active:    "inactive",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return
}

func (c drawerImp) convertToBSONUpdate(payload requestmodel.DrawerUpdate, doc docmodel.Drawer) (result docmodel.Drawer) {
	result = docmodel.Drawer{
		ID:        doc.ID,
		Name:      payload.Name,
		Cabinet:   payload.Cabinet,
		Active:    doc.Active,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: time.Now(),
	}
	return
}

func (c drawerImp) convertToBSONUpdateActive(payload requestmodel.DrawerUpdateActive, doc docmodel.Drawer) (result docmodel.Drawer) {
	result = docmodel.Drawer{
		ID:        doc.ID,
		Name:      doc.Name,
		Cabinet:   doc.Cabinet,
		Active:    payload.Active,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: time.Now(),
	}
	return
}

func (c drawerImp) convertToJSON(ctx context.Context, doc docmodel.Drawer) (res responsemodel.DrawerTable) {

	cabinet, _ := Cabinet().FindOneWithId(ctx, doc.Cabinet)

	res = responsemodel.DrawerTable{
		ID:        doc.ID.Hex(),
		Cabinet:   cabinet,
		Name:      doc.Name,
		Active:    doc.Active,
		CreatedAt: doc.CreatedAt.String(),
	}
	return
}
