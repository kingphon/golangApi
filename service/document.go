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

type DocumentInterface interface {
	Create(ctx context.Context, payload requestmodel.DocumentCreate) (err error)

	All(ctx context.Context, query requestmodel.DocumentAll) (result responsemodel.DocumentAll)

	UpdateActive(ctx context.Context, payload requestmodel.DocumentUpdateStatus) (err error)

	Update(ctx context.Context, payload requestmodel.DocumentUpdate) (err error)

	FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.DocumentTable, err error)
}

type documentImp struct {
}

func Document() DocumentInterface {
	return documentImp{}
}

func (c documentImp) Create(ctx context.Context, payload requestmodel.DocumentCreate) (err error) {

	count, err := dao.Drawer().Count(ctx, bson.M{"_id": payload.Drawer})

	if count == 0 {
		err = errors.New("kệ không tồn tại")
		return
	}

	oid, _ := primitive.ObjectIDFromHex(payload.Drawer)

	count, err = dao.Company().Count(ctx, bson.M{"_id": oid})

	if count == 0 {
		err = errors.New("công ty không tồn tại")
		return
	}

	docBSON := c.convertToBSON(payload, oid)

	err = dao.Document().Create(ctx, docBSON)

	if err != nil {
		err = errors.New("không thể tạo ngăn")
	}

	return
}

func (c documentImp) UpdateActive(ctx context.Context, payload requestmodel.DocumentUpdateStatus) (err error) {

	doc, err := dao.Document().FindOne(ctx, bson.M{"_id": payload.ID})

	if err != nil {
		return
	}

	docBSON := c.convertToBSONUpdateActive(payload, doc)

	err = dao.Document().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật ngăn")
	}

	return
}

func (c documentImp) Update(ctx context.Context, payload requestmodel.DocumentUpdate) (err error) {

	doc, err := dao.Document().FindOne(ctx, bson.D{{"_id", payload.ID}})

	if err != nil {
		return
	}

	oid, _ := primitive.ObjectIDFromHex(payload.Drawer)

	count, err := dao.Company().Count(ctx, bson.M{"_id": oid})

	if count == 0 {
		err = errors.New("công ty không tồn tại")
		return
	}

	docBSON := c.convertToBSONUpdate(payload, doc, oid)

	err = dao.Document().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật ngăn")
	}

	return
}

func (c documentImp) All(ctx context.Context, query requestmodel.DocumentAll) (result responsemodel.DocumentAll) {

	opts := options.FindOptions{}
	opts.SetLimit(query.Limit)
	opts.SetSkip(query.Page)

	docs := dao.Document().Find(ctx, bson.M{}, &opts)

	for _, doc := range docs {
		result.Data = append(result.Data, c.convertToJSON(ctx, doc))
	}

	total, err := dao.Document().Count(ctx, bson.M{})
	if err != nil {
		err = errors.New("đã xảy ra lỗi")
	}

	result.Total = total
	result.Limit = query.Limit

	return

}

func (c documentImp) FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.DocumentTable, err error) {

	doc, err := dao.Document().FindOne(ctx, bson.M{"_id": id})

	if err != nil {
		err = errors.New("không thể tạo ngăn")
	}

	result = c.convertToJSON(ctx, doc)

	return
}

func (c documentImp) convertToBSON(payload requestmodel.DocumentCreate, drawerId primitive.ObjectID) (doc docmodel.Document) {
	doc = docmodel.Document{
		ID:        primitive.NewObjectID(),
		Title:     payload.Title,
		Content:   payload.Content,
		Drawer:    drawerId,
		Status:    "waiting_for_verify",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return
}

func (c documentImp) convertToBSONUpdate(payload requestmodel.DocumentUpdate, doc docmodel.Document, drawerId primitive.ObjectID) (result docmodel.Document) {
	result = docmodel.Document{
		ID:        doc.ID,
		Title:     payload.Title,
		Content:   payload.Content,
		Drawer:    drawerId,
		Status:    doc.Status,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: time.Now(),
	}
	return
}

func (c documentImp) convertToBSONUpdateActive(payload requestmodel.DocumentUpdateStatus, doc docmodel.Document) (result docmodel.Document) {
	result = docmodel.Document{
		ID:        doc.ID,
		Title:     doc.Title,
		Content:   doc.Content,
		Drawer:    doc.Drawer,
		Status:    payload.Status,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: time.Now(),
	}
	return
}

func (c documentImp) convertToJSON(ctx context.Context, doc docmodel.Document) (res responsemodel.DocumentTable) {

	drawer, _ := Drawer().FindOneWithId(ctx, doc.Drawer)

	res = responsemodel.DocumentTable{
		ID:        doc.ID.Hex(),
		Drawer:    drawer,
		Title:     doc.Title,
		Content:   doc.Content,
		Status:    doc.Status,
		CreatedAt: doc.CreatedAt.String(),
	}
	return
}
