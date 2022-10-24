package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	DrawerCreate struct {
		Name          string `json:"name"`
		CabinetString string `json:"cabinet"`
		Cabinet       primitive.ObjectID
	}

	DrawerUpdate struct {
		ID            primitive.ObjectID
		Name          string `json:"name"`
		CabinetString string `json:"cabinet"`
		Cabinet       primitive.ObjectID
	}

	DrawerUpdateActive struct {
		ID     primitive.ObjectID `query:"id""`
		Active string             `json:"active"`
	}

	DrawerAll struct {
		Page  int64 `query:"page"`
		Limit int64 `query:"limit"`
	}

	DrawerDetail struct {
		ID primitive.ObjectID
	}
)

// Validate ...
func (c DrawerCreate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required.Error("tên không được trống")),
		validation.Field(&c.CabinetString,
			validation.Required.Error("kệ không được trống")),
	)
}

func (c DrawerUpdate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required.Error("tên không được trống")),
		validation.Field(&c.Cabinet,
			validation.Required.Error("kệ không được trống")),
	)
}

func (c DrawerUpdateActive) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Active,
			validation.Required.Error("trạng thái không được trống"), validation.In("active", "inactive").Error("trạng thái không tồn tại")),
	)
}
