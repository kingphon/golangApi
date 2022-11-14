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

type DepartmentInterface interface {
	Create(ctx context.Context, payload requestmodel.DepartmentCreate) (err error)

	All(ctx context.Context, query requestmodel.DepartmentAll) (result responsemodel.DepartmentAll)

	UpdateActive(ctx context.Context, payload requestmodel.DepartmentUpdateActive, id primitive.ObjectID) (err error)

	Update(ctx context.Context, payload requestmodel.DepartmentUpdate, id primitive.ObjectID) (err error)

	FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.DepartmentTable, err error)
}

type departmentImp struct {
}

func Department() DepartmentInterface {
	return departmentImp{}
}

func (c departmentImp) Create(ctx context.Context, payload requestmodel.DepartmentCreate) (err error) {

	var permissionArray []string

	oid, _ := primitive.ObjectIDFromHex(payload.Company)

	count, err := dao.Company().Count(ctx, bson.M{"_id": oid})

	if count == 0 {
		err = errors.New("công ty không tồn tại")
		return
	}

	for _, per := range payload.Permission {
		if count, _ := dao.Permission().Count(ctx, bson.M{"code": per}); count != 0 {
			permissionArray = append(permissionArray, per)
		}
	}

	payload.Permission = permissionArray

	docBSON := c.convertToBSON(payload, oid)

	err = dao.Department().Create(ctx, docBSON)

	if err != nil {
		err = errors.New("không thể tạo phòng ban")
	}

	return
}

func (c departmentImp) UpdateActive(ctx context.Context, payload requestmodel.DepartmentUpdateActive, id primitive.ObjectID) (err error) {

	doc, err := dao.Department().FindOne(ctx, bson.M{"_id": id})

	if err != nil {
		return
	}

	docBSON := c.convertToBSONUpdateActive(payload, doc)

	err = dao.Department().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật phòng ban")
	}

	return
}

func (c departmentImp) Update(ctx context.Context, payload requestmodel.DepartmentUpdate, id primitive.ObjectID) (err error) {

	doc, err := dao.Department().FindOne(ctx, bson.D{{"_id", id}})

	if err != nil {
		return
	}

	var permissionArray []string

	oid, _ := primitive.ObjectIDFromHex(payload.Company)

	count, err := dao.Company().Count(ctx, bson.M{"_id": oid})

	if count == 0 {
		err = errors.New("công ty không tồn tại")
		return
	}

	for _, per := range payload.Permission {
		if count, _ := dao.Permission().Count(ctx, bson.M{"code": per}); count != 0 {
			permissionArray = append(permissionArray, per)
		}
	}

	payload.Permission = permissionArray

	docBSON := c.convertToBSONUpdate(payload, doc, oid)

	err = dao.Department().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật phòng ban")
	}

	return
}

func (c departmentImp) All(ctx context.Context, query requestmodel.DepartmentAll) (result responsemodel.DepartmentAll) {

	opts := options.FindOptions{}
	opts.SetLimit(query.Limit)
	opts.SetSkip(query.Page)

	docs := dao.Department().Find(ctx, bson.M{}, &opts)

	for _, doc := range docs {
		result.Data = append(result.Data, c.convertToJSON(ctx, doc))
	}

	total, err := dao.Department().Count(ctx, bson.M{})
	if err != nil {
		err = errors.New("đã xảy ra lỗi")
	}

	result.Total = total
	result.Limit = query.Limit

	return

}

func (c departmentImp) FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.DepartmentTable, err error) {

	doc, err := dao.Department().FindOne(ctx, bson.M{"_id": id})

	if err != nil {
		err = errors.New("không thể tìm thấy phòng ban")
	}

	result = c.convertToJSON(ctx, doc)

	return
}

func (c departmentImp) convertToBSON(payload requestmodel.DepartmentCreate, companyId primitive.ObjectID) (doc docmodel.Department) {
	doc = docmodel.Department{
		ID:         primitive.NewObjectID(),
		Name:       payload.Name,
		Company:    companyId,
		Permission: payload.Permission,
		Active:     "inactive",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	return
}

func (c departmentImp) convertToBSONUpdate(payload requestmodel.DepartmentUpdate, doc docmodel.Department, companyId primitive.ObjectID) (result docmodel.Department) {
	result = docmodel.Department{
		ID:         doc.ID,
		Name:       payload.Name,
		Permission: payload.Permission,
		Company:    companyId,
		Active:     doc.Active,
		CreatedAt:  doc.CreatedAt,
		UpdatedAt:  time.Now(),
	}
	return
}

func (c departmentImp) convertToBSONUpdateActive(payload requestmodel.DepartmentUpdateActive, doc docmodel.Department) (result docmodel.Department) {
	result = docmodel.Department{
		ID:         doc.ID,
		Name:       doc.Name,
		Company:    doc.Company,
		Active:     payload.Active,
		Permission: doc.Permission,
		CreatedAt:  doc.CreatedAt,
		UpdatedAt:  time.Now(),
	}
	return
}

func (c departmentImp) convertToJSON(ctx context.Context, doc docmodel.Department) (res responsemodel.DepartmentTable) {

	company, _ := Company().FindOneWithId(ctx, doc.Company)

	res = responsemodel.DepartmentTable{
		ID:         doc.ID.Hex(),
		Company:    company,
		Name:       doc.Name,
		Permission: doc.Permission,
		Active:     doc.Active,
		CreatedAt:  doc.CreatedAt.String(),
	}
	return
}
