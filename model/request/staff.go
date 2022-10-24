package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	StaffCreate struct {
		Name       string `json:"name"`
		Phone      string `json:"phone"`
		Password   string `json:"password"`
		Department string `json:"department"`
	}

	StaffUpdate struct {
		Name       string `json:"name"`
		Phone      string `json:"phone"`
		Department string `json:"department"`
	}

	StaffUpdateActive struct {
		Active string `json:"active"`
	}

	StaffUpdatePassword struct {
		OldPassword        string `json:"oldPassword"`
		ConfirmOldPassword string `json:"confirmOldPassword"`
		Password           string `json:"password"`
	}

	StaffLogin struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	StaffAll struct {
		Page  int64 `query:"page"`
		Limit int64 `query:"limit"`
	}
)

// Validate ...
func (s StaffCreate) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Name,
			validation.Required.Error("tên không được trống")),
		validation.Field(&s.Department,
			validation.Required.Error("phòng ban không được trống")),
		validation.Field(&s.Phone,
			validation.Required.Error("số điện thoại không được trống")),
		validation.Field(&s.Password,
			validation.Required.Error("mật khẩu không được trống")),
	)
}

func (s StaffUpdate) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Name,
			validation.Required.Error("tên không được trống")),
		validation.Field(&s.Department,
			validation.Required.Error("phòng ban không được trống")),
		validation.Field(&s.Phone,
			validation.Required.Error("số điện thoại không được trống")),
	)
}

func (s StaffUpdateActive) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Active,
			validation.Required.Error("trạng thái không được trống"), validation.In("active", "inactive").Error("trạng thái không tồn tại")),
	)
}

func (s StaffUpdatePassword) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Password,
			validation.Required.Error("mật khẩu mới không được trống")),
		validation.Field(&s.OldPassword,
			validation.Required.Error("mật khẩu cũ không được trống")),
		validation.Field(&s.ConfirmOldPassword,
			validation.Required.Error("xác nhận mật khẩu cũ không được trống")),
	)
}

func (s StaffLogin) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Password,
			validation.Required.Error("mật khẩu mới không được trống")),
		validation.Field(&s.Phone,
			validation.Required.Error("số điện thoại không được trống")),
	)
}
