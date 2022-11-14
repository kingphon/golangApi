package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	DocumentCreate struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Drawer  string `json:"drawer"`
	}

	DocumentUpdate struct {
		ID      primitive.ObjectID
		Title   string `json:"title"`
		Content string `json:"content"`
		Drawer  string `json:"drawer"`
	}

	DocumentUpdateStatus struct {
		ID     primitive.ObjectID `query:"id""`
		Status string             `json:"status"`
	}

	DocumentAll struct {
		Page  int64 `query:"page"`
		Limit int64 `query:"limit"`
	}

	DocumentDetail struct {
		ID primitive.ObjectID
	}
)

// Validate ...
func (c DocumentCreate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title,
			validation.Required.Error("tiêu đề không được trống")),
		validation.Field(&c.Drawer,
			validation.Required.Error("ngăn không được trống")),
		validation.Field(&c.Content,
			validation.Required.Error("nội dung không được trống")),
	)
}

func (c DocumentUpdate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title,
			validation.Required.Error("tiêu đề không được trống")),
		validation.Field(&c.Drawer,
			validation.Required.Error("ngăn không được trống")),
		validation.Field(&c.Content,
			validation.Required.Error("nội dung không được trống")),
	)
}

func (c DocumentUpdateStatus) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Status,
			validation.Required.Error("trạng thái không được trống"), validation.In("waiting_for_verify", "verified", "delivered").Error("trạng thái không tồn tại")),
	)
}
