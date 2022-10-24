package responsemodel

type (
	CabinetCreate struct {
		ID string `json:"_id"`
	}

	CabinetUpdate struct {
		ID string `json:"_id"`
	}

	CabinetUpdateStatus struct {
		ID string `json:"_id"`
	}

	CabinetAll struct {
		Limit int64          `json:"limit"`
		Total int64          `json:"total"`
		Data  []CabinetTable `json:"data"`
	}

	CabinetTable struct {
		ID        string       `json:"_id"`
		Company   CompanyTable `json:"company"`
		Name      string       `json:"name"`
		Active    string       `json:"active"`
		CreatedAt string       `json:"createdAt"`
	}
)
