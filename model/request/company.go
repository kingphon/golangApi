package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	CompanyCreate struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	CompanyUpdate struct {
		Name string `json:"name"`
	}

	CompanyUpdateActive struct {
		Active string `json:"active"`
	}

	CompanyAll struct {
		Page  int64 `query:"page"`
		Limit int64 `query:"limit"`
	}
)

// Validate ...
func (c CompanyCreate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required.Error("tên không được trống")),
		validation.Field(&c.Type,
			validation.Required.Error("loại không được trống")),
	)
}

func (c CompanyUpdate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required.Error("tên không được trống")),
		//validation.Field(&c.Type,
		//	validation.Required.Error("loại không được trống")),
	)
}

func (c CompanyUpdateActive) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Active,
			validation.Required.Error("trạng thái không được trống"), validation.In("active", "inactive").Error("trạng thái không tồn tại")),
	)
}
