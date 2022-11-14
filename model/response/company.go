package responsemodel

type (
	CompanyCreate struct {
		ID string `json:"_id"`
	}

	CompanyUpdate struct {
		ID string `json:"_id"`
	}

	CompanyUpdateStatus struct {
		ID string `json:"_id"`
	}

	CompanyAll struct {
		Limit int64          `json:"limit"`
		Total int64          `json:"total"`
		Data  []CompanyTable `json:"data"`
	}

	CompanyTable struct {
		ID        string           `json:"_id"`
		Type      CompanyTypeTable `json:"type"`
		Name      string           `json:"name"`
		Active    string           `json:"active"`
		CreatedAt string           `json:"createdAt"`
	}
)
