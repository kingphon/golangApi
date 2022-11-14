package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	PermissionCreate struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	PermissionUpdate struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	PermissionAll struct {
		Page  int64 `query:"page"`
		Limit int64 `query:"limit"`
	}
)

// Validate ...
func (c PermissionCreate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required.Error("tên không được trống")),
		validation.Field(&c.Code,
			validation.Required.Error("mã không được trống")),
	)
}

func (c PermissionUpdate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required.Error("tên không được trống")),
		validation.Field(&c.Code,
			validation.Required.Error("mã không được trống")),
	)
}
