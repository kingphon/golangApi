package requestmodel

import validation "github.com/go-ozzo/ozzo-validation"

type (
	CompanyTypeCreate struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	CompanyTypeAll struct {
		Page  int64 `query:"page"`
		Limit int64 `query:"limit"`
	}
)

// Validate ...
func (ct CompanyTypeCreate) Validate() error {
	return validation.ValidateStruct(&ct,
		validation.Field(&ct.Name,
			validation.Required.Error("tên không được trống")),
		validation.Field(&ct.Code,
			validation.Required.Error("code không được trống")),
	)
}

// Validate ...
func (ct CompanyTypeAll) Validate() error {
	return validation.ValidateStruct(&ct)
}
