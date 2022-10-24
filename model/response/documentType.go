package responsemodel

type (
	DocumentTypeCreate struct {
		ID string `json:"_id"`
	}

	DocumentTypeAll struct {
		Limit int64               `json:"limit"`
		Total int64               `json:"total"`
		Data  []DocumentTypeTable `json:"data"`
	}

	DocumentTypeTable struct {
		ID   string `json:"_id"`
		Code string `json:"code"`
		Name string `json:"name"`
	}
)
