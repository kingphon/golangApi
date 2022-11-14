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
	"golangApi/util"
	"time"
)

type StaffInterface interface {
	Create(ctx context.Context, payload requestmodel.StaffCreate) (err error)

	All(ctx context.Context, query requestmodel.StaffAll) (result responsemodel.StaffAll)

	UpdateActive(ctx context.Context, payload requestmodel.StaffUpdateActive, id primitive.ObjectID) (err error)

	Update(ctx context.Context, payload requestmodel.StaffUpdate, id primitive.ObjectID) (err error)

	FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.StaffTable, err error)

	UpdatePassword(ctx context.Context, payload requestmodel.StaffUpdatePassword, id primitive.ObjectID) (err error)

	Login(ctx context.Context, payload requestmodel.StaffLogin) (token string, err error)
}

type staffImp struct {
}

func Staff() StaffInterface {
	return staffImp{}
}

func (s staffImp) Create(ctx context.Context, payload requestmodel.StaffCreate) (err error) {

	oid, _ := primitive.ObjectIDFromHex(payload.Department)

	count, err := dao.Department().Count(ctx, bson.M{"_id": oid})

	if count == 0 {
		err = errors.New("nhân viên không tồn tại")
		return
	}

	docBSON := s.convertToBSON(payload, oid)

	err = dao.Staff().Create(ctx, docBSON)

	if err != nil {
		err = errors.New("không thể tạo nhân viên")
	}

	return
}

func (s staffImp) UpdateActive(ctx context.Context, payload requestmodel.StaffUpdateActive, id primitive.ObjectID) (err error) {

	doc, err := dao.Staff().FindOne(ctx, bson.M{"_id": id})

	if err != nil {
		return
	}

	docBSON := s.convertToBSONUpdateActive(payload, doc)

	err = dao.Staff().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật nhân viên")
	}

	return
}

func (s staffImp) Update(ctx context.Context, payload requestmodel.StaffUpdate, id primitive.ObjectID) (err error) {

	oid, _ := primitive.ObjectIDFromHex(payload.Department)

	count, err := dao.Department().Count(ctx, bson.M{"_id": oid})

	if count == 0 {
		err = errors.New("nhân viên không tồn tại")
		return
	}

	doc, err := dao.Staff().FindOne(ctx, bson.D{{"_id", id}})

	if err != nil {
		return
	}

	docBSON := s.convertToBSONUpdate(payload, doc, oid)

	err = dao.Staff().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật nhân viên")
	}

	return
}

func (s staffImp) All(ctx context.Context, query requestmodel.StaffAll) (result responsemodel.StaffAll) {

	opts := options.FindOptions{}
	opts.SetLimit(query.Limit)
	opts.SetSkip(query.Page)

	docs := dao.Staff().Find(ctx, bson.M{}, &opts)

	for _, doc := range docs {
		result.Data = append(result.Data, s.convertToJSON(ctx, doc))
	}

	total, err := dao.Staff().Count(ctx, bson.M{})
	if err != nil {
		err = errors.New("đã xảy ra lỗi")
	}

	result.Total = total
	result.Limit = query.Limit

	return

}

func (s staffImp) FindOneWithId(ctx context.Context, id primitive.ObjectID) (result responsemodel.StaffTable, err error) {

	doc, err := dao.Staff().FindOne(ctx, bson.M{"_id": id})

	if err != nil {
		err = errors.New("không thể tìm thấy nhân viên")
	}

	result = s.convertToJSON(ctx, doc)

	return
}

func (s staffImp) UpdatePassword(ctx context.Context, payload requestmodel.StaffUpdatePassword, id primitive.ObjectID) (err error) {

	if payload.ConfirmOldPassword != payload.OldPassword {
		err = errors.New("xác nhận mật khẩu không chính xác")
		return
	}

	doc, err := dao.Staff().FindOne(ctx, bson.M{"_id": id})

	if err != nil {
		return
	}

	match := util.CheckPasswordHash(payload.OldPassword, doc.Password)

	if match == false {
		err = errors.New("mật khẩu cũ không chính xác")
		return
	}

	docBSON := s.convertToBSONUpdatePassword(payload, doc)

	err = dao.Staff().UpdateOne(ctx, bson.M{"_id": docBSON.ID}, docBSON)

	if err != nil {
		err = errors.New("không thể cập nhật nhân viên")
	}

	return
}

func (s staffImp) Login(ctx context.Context, payload requestmodel.StaffLogin) (token string, err error) {

	doc, err := dao.Staff().FindOne(ctx, bson.M{"phone": payload.Phone})

	if err != nil {
		return
	}

	match := util.CheckPasswordHash(payload.Password, doc.Password)

	if match == false {
		err = errors.New("đăng nhập không thành công")
		return
	}

	data := s.convertToJSON(ctx, doc)

	token, err = util.GenerateToken(data)

	return
}

func (s staffImp) convertToBSON(payload requestmodel.StaffCreate, departmentId primitive.ObjectID) (doc docmodel.Staff) {
	hash, _ := util.HashPassword(payload.Password)
	doc = docmodel.Staff{
		ID:         primitive.NewObjectID(),
		Name:       payload.Name,
		Phone:      payload.Phone,
		Password:   hash,
		Department: departmentId,
		Active:     "inactive",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	return
}

func (s staffImp) convertToBSONUpdate(payload requestmodel.StaffUpdate, doc docmodel.Staff, departmentId primitive.ObjectID) (result docmodel.Staff) {
	result = docmodel.Staff{
		ID:         doc.ID,
		Name:       payload.Name,
		Phone:      payload.Phone,
		Password:   doc.Password,
		Department: departmentId,
		Active:     doc.Active,
		CreatedAt:  doc.CreatedAt,
		UpdatedAt:  time.Now(),
	}
	return
}

func (s staffImp) convertToBSONUpdateActive(payload requestmodel.StaffUpdateActive, doc docmodel.Staff) (result docmodel.Staff) {
	result = docmodel.Staff{
		ID:         doc.ID,
		Name:       doc.Name,
		Phone:      doc.Phone,
		Password:   doc.Password,
		Department: doc.Department,
		Active:     payload.Active,
		CreatedAt:  doc.CreatedAt,
		UpdatedAt:  time.Now(),
		IsRoot:     false,
	}
	return
}

func (s staffImp) convertToBSONUpdatePassword(payload requestmodel.StaffUpdatePassword, doc docmodel.Staff) (result docmodel.Staff) {
	hash, _ := util.HashPassword(payload.Password)
	result = docmodel.Staff{
		ID:         doc.ID,
		Name:       doc.Name,
		Phone:      doc.Phone,
		Password:   hash,
		Department: doc.Department,
		Active:     doc.Active,
		CreatedAt:  doc.CreatedAt,
		UpdatedAt:  time.Now(),
		IsRoot:     doc.IsRoot,
	}
	return
}

func (s staffImp) convertToJSON(ctx context.Context, doc docmodel.Staff) (res responsemodel.StaffTable) {

	department, _ := Department().FindOneWithId(ctx, doc.Department)

	res = responsemodel.StaffTable{
		ID:         doc.ID.Hex(),
		Department: department,
		Phone:      doc.Phone,
		Name:       doc.Name,
		IsRoot:     doc.IsRoot,
		Active:     doc.Active,
		CreatedAt:  doc.CreatedAt.String(),
	}
	return
}
