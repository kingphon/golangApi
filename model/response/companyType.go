package responsemodel

type (
	CompanyTypeCreate struct {
		ID string `json:"_id"`
	}

	CompanyTypeAll struct {
		Limit int64              `json:"limit"`
		Total int64              `json:"total"`
		Data  []CompanyTypeTable `json:"data"`
	}

	CompanyTypeTable struct {
		ID   string `json:"_id"`
		Code string `json:"code"`
		Name string `json:"name"`
	}
)
