package responsemodel

type (
	DepartmentCreate struct {
		ID string `json:"_id"`
	}

	DepartmentUpdate struct {
		ID string `json:"_id"`
	}

	DepartmentUpdateStatus struct {
		ID string `json:"_id"`
	}

	DepartmentAll struct {
		Limit int64             `json:"limit"`
		Total int64             `json:"total"`
		Data  []DepartmentTable `json:"data"`
	}

	DepartmentTable struct {
		ID         string       `json:"_id"`
		Company    CompanyTable `json:"company"`
		Name       string       `json:"name"`
		Permission []string     `json:"permission"`
		Active     string       `json:"active"`
		CreatedAt  string       `json:"createdAt"`
	}
)
