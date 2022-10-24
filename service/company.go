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

type CompanyInterface interface {
	Create(ctx context.Context, payload requestmodel.CompanyCreate) (err error)

	All(ctx context.Context, query requestmodel.CompanyAll) (result responsemodel.CompanyAll)

	UpdateActive(ctx context.Context, payload requestmodel.CompanyUpdateActive, id primitive.ObjectID) (err error)

	Update(ctx context.Context, payload requestmodel.CompanyUpdate, id primitive.ObjectID) (err error)

	FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.CompanyTable, err error)
}

type companyImp struct {
}

func Company() CompanyInterface {
	return companyImp{}
}

func (c companyImp) Create(ctx context.Context, payload requestmodel.CompanyCreate) (err error) {

	oid, _ := primitive.ObjectIDFromHex(payload.Type)

	count, err := dao.CompanyType().Count(ctx, bson.M{"_id": oid})

	if count == 0 {
		err = errors.New("loại công ty không tồn tại")
		return
	}

	docBSON := c.convertToBSON(payload, oid)

	err = dao.Company().Create(ctx, docBSON)

	if err != nil {
		err = errors.New("không thể tạo công ty")
	}

	return
}

func (c companyImp) UpdateActive(ctx context.Context, payload requestmodel.CompanyUpdateActive, id primitive.ObjectID) (err error) {

	doc, err := dao.Company().FindOne(ctx, bson.M{"_id": id})

	if err != nil {
		return
	}

	docBSON := c.convertToBSONUpdateActive(payload, doc)

	err = dao.Company().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật công ty")
	}

	return
}

func (c companyImp) Update(ctx context.Context, payload requestmodel.CompanyUpdate, id primitive.ObjectID) (err error) {

	doc, err := dao.Company().FindOne(ctx, bson.D{{"_id", id}})

	if err != nil {
		return
	}

	docBSON := c.convertToBSONUpdate(payload, doc)

	err = dao.Company().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật công ty")
	}

	return
}

func (c companyImp) All(ctx context.Context, query requestmodel.CompanyAll) (result responsemodel.CompanyAll) {

	opts := options.FindOptions{}
	opts.SetLimit(query.Limit)
	opts.SetSkip(query.Page)

	docs := dao.Company().Find(ctx, bson.M{}, &opts)

	for _, doc := range docs {
		result.Data = append(result.Data, c.convertToJSON(ctx, doc))
	}

	total, err := dao.Company().Count(ctx, bson.M{})
	if err != nil {
		err = errors.New("đã xảy ra lỗi")
	}

	result.Total = total
	result.Limit = query.Limit

	return

}

func (c companyImp) FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.CompanyTable, err error) {

	doc, err := dao.Company().FindOne(ctx, bson.M{"_id": id})

	if err != nil {
		err = errors.New("không thể tạo công ty")
	}

	result = c.convertToJSON(ctx, doc)

	return
}

func (c companyImp) convertToBSON(payload requestmodel.CompanyCreate, typeId primitive.ObjectID) (doc docmodel.Company) {
	doc = docmodel.Company{
		ID:        primitive.NewObjectID(),
		Name:      payload.Name,
		Type:      typeId,
		Active:    "inactive",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return
}

func (c companyImp) convertToBSONUpdate(payload requestmodel.CompanyUpdate, doc docmodel.Company) (result docmodel.Company) {
	result = docmodel.Company{
		ID:        doc.ID,
		Name:      payload.Name,
		Type:      doc.Type,
		Active:    doc.Active,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: time.Now(),
	}
	return
}

func (c companyImp) convertToBSONUpdateActive(payload requestmodel.CompanyUpdateActive, doc docmodel.Company) (result docmodel.Company) {
	result = docmodel.Company{
		ID:        doc.ID,
		Name:      doc.Name,
		Type:      doc.Type,
		Active:    payload.Active,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: time.Now(),
	}
	return
}

func (c companyImp) convertToJSON(ctx context.Context, doc docmodel.Company) (res responsemodel.CompanyTable) {

	companyType, _ := CompanyType().FindOneWithId(ctx, doc.Type)

	res = responsemodel.CompanyTable{
		ID:        doc.ID.Hex(),
		Type:      companyType,
		Name:      doc.Name,
		Active:    doc.Active,
		CreatedAt: doc.CreatedAt.String(),
	}
	return
}
