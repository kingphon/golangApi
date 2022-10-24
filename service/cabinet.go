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

type CabinetInterface interface {
	Create(ctx context.Context, payload requestmodel.CabinetCreate) (err error)

	All(ctx context.Context, query requestmodel.CabinetAll) (result responsemodel.CabinetAll)

	UpdateActive(ctx context.Context, payload requestmodel.CabinetUpdateActive) (err error)

	Update(ctx context.Context, payload requestmodel.CabinetUpdate) (err error)

	FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.CabinetTable, err error)
}

type cabinetImp struct {
}

func Cabinet() CabinetInterface {
	return cabinetImp{}
}

func (c cabinetImp) Create(ctx context.Context, payload requestmodel.CabinetCreate) (err error) {

	count, err := dao.Company().Count(ctx, bson.M{"_id": payload.Company})

	if count == 0 {
		err = errors.New("công ty không tồn tại")
		return
	}

	docBSON := c.convertToBSON(payload)

	err = dao.Cabinet().Create(ctx, docBSON)

	if err != nil {
		err = errors.New("không thể tạo kệ")
	}

	return
}

func (c cabinetImp) UpdateActive(ctx context.Context, payload requestmodel.CabinetUpdateActive) (err error) {

	doc, err := dao.Cabinet().FindOne(ctx, bson.M{"_id": payload.ID})

	if err != nil {
		return
	}

	docBSON := c.convertToBSONUpdateActive(payload, doc)

	err = dao.Cabinet().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật kệ")
	}

	return
}

func (c cabinetImp) Update(ctx context.Context, payload requestmodel.CabinetUpdate) (err error) {

	doc, err := dao.Cabinet().FindOne(ctx, bson.D{{"_id", payload.ID}})

	if err != nil {
		return
	}

	docBSON := c.convertToBSONUpdate(payload, doc)

	err = dao.Cabinet().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật kệ")
	}

	return
}

func (c cabinetImp) All(ctx context.Context, query requestmodel.CabinetAll) (result responsemodel.CabinetAll) {

	opts := options.FindOptions{}
	opts.SetLimit(query.Limit)
	opts.SetSkip(query.Page)

	docs := dao.Cabinet().Find(ctx, bson.M{}, &opts)

	for _, doc := range docs {
		result.Data = append(result.Data, c.convertToJSON(ctx, doc))
	}

	total, err := dao.Cabinet().Count(ctx, bson.M{})
	if err != nil {
		err = errors.New("đã xảy ra lỗi")
	}

	result.Total = total
	result.Limit = query.Limit

	return

}

func (c cabinetImp) FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.CabinetTable, err error) {

	doc, err := dao.Cabinet().FindOne(ctx, bson.M{"_id": id})

	if err != nil {
		err = errors.New("đã có lỗi xảy ra")
	}

	result = c.convertToJSON(ctx, doc)

	return
}

func (c cabinetImp) convertToBSON(payload requestmodel.CabinetCreate) (doc docmodel.Cabinet) {
	doc = docmodel.Cabinet{
		ID:        primitive.NewObjectID(),
		Name:      payload.Name,
		Company:   payload.Company,
		Active:    "inactive",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return
}

func (c cabinetImp) convertToBSONUpdate(payload requestmodel.CabinetUpdate, doc docmodel.Cabinet) (result docmodel.Cabinet) {
	result = docmodel.Cabinet{
		ID:        doc.ID,
		Name:      payload.Name,
		Company:   payload.Company,
		Active:    doc.Active,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: time.Now(),
	}
	return
}

func (c cabinetImp) convertToBSONUpdateActive(payload requestmodel.CabinetUpdateActive, doc docmodel.Cabinet) (result docmodel.Cabinet) {
	result = docmodel.Cabinet{
		ID:        doc.ID,
		Name:      doc.Name,
		Company:   doc.Company,
		Active:    payload.Active,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: time.Now(),
	}
	return
}

func (c cabinetImp) convertToJSON(ctx context.Context, doc docmodel.Cabinet) (res responsemodel.CabinetTable) {

	company, _ := Company().FindOneWithId(ctx, doc.Company)

	res = responsemodel.CabinetTable{
		ID:        doc.ID.Hex(),
		Company:   company,
		Name:      doc.Name,
		Active:    doc.Active,
		CreatedAt: doc.CreatedAt.String(),
	}
	return
}
