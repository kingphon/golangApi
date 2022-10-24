package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CabinetCreate struct {
		Name          string `json:"name"`
		CompanyString string `json:"company"`
		Company       primitive.ObjectID
	}

	CabinetUpdate struct {
		ID            primitive.ObjectID
		Name          string `json:"name"`
		CompanyString string `json:"company"`
		Company       primitive.ObjectID
	}

	CabinetUpdateActive struct {
		ID     primitive.ObjectID `query:"id""`
		Active string             `json:"active"`
	}

	CabinetAll struct {
		Page  int64 `query:"page"`
		Limit int64 `query:"limit"`
	}

	CabinetDetail struct {
		ID primitive.ObjectID
	}
)

// Validate ...
func (c CabinetCreate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required.Error("tên không được trống")),
		validation.Field(&c.CompanyString,
			validation.Required.Error("công ty không được trống")),
	)
}

func (c CabinetUpdate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required.Error("tên không được trống")),
		validation.Field(&c.CompanyString,
			validation.Required.Error("công ty không được trống")),
	)
}

func (c CabinetUpdateActive) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Active,
			validation.Required.Error("trạng thái không được trống"), validation.In("active", "inactive").Error("trạng thái không tồn tại")),
	)
}
