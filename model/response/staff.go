package responsemodel

type (
	StaffCreate struct {
		ID string `json:"_id"`
	}

	StaffUpdate struct {
		ID string `json:"_id"`
	}

	StaffUpdateStatus struct {
		ID string `json:"_id"`
	}

	StaffAll struct {
		Limit int64        `json:"limit"`
		Total int64        `json:"total"`
		Data  []StaffTable `json:"data"`
	}

	StaffTable struct {
		ID         string          `json:"_id"`
		Department DepartmentTable `json:"department"`
		Name       string          `json:"name"`
		Phone      string          `json:"phone"`
		IsRoot     bool            `json:"isRoot"`
		Active     string          `json:"active"`
		CreatedAt  string          `json:"createdAt"`
	}
)
