package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	DepartmentCreate struct {
		Name       string   `json:"name"`
		Company    string   `json:"company"`
		Permission []string `json:"permission"`
	}

	DepartmentUpdate struct {
		Name       string   `json:"name"`
		Company    string   `json:"company"`
		Permission []string `json:"permission"`
	}

	DepartmentUpdateActive struct {
		Active string `json:"active"`
	}

	DepartmentAll struct {
		Page  int64 `query:"page"`
		Limit int64 `query:"limit"`
	}
)

// Validate ...
func (d DepartmentCreate) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name,
			validation.Required.Error("tên không được trống")),
		validation.Field(&d.Company,
			validation.Required.Error("công ty không được trống")),
	)
}

func (d DepartmentUpdate) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name,
			validation.Required.Error("tên không được trống")),
		validation.Field(&d.Company,
			validation.Required.Error("công ty không được trống")),
	)
}

func (d DepartmentUpdateActive) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Active,
			validation.Required.Error("trạng thái không được trống"), validation.In("active", "inactive").Error("trạng thái không tồn tại")),
	)
}
