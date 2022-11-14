package responsemodel

type (
	DocumentCreate struct {
		ID string `json:"_id"`
	}

	DocumentUpdate struct {
		ID string `json:"_id"`
	}

	DocumentUpdateStatus struct {
		ID string `json:"_id"`
	}

	DocumentAll struct {
		Limit int64           `json:"limit"`
		Total int64           `json:"total"`
		Data  []DocumentTable `json:"data"`
	}

	DocumentTable struct {
		ID        string      `json:"_id"`
		Drawer    DrawerTable `json:"drawer"`
		Title     string      `json:"title"`
		Content   string      `json:"content"`
		Status    string      `json:"status"`
		CreatedAt string      `json:"createdAt"`
	}
)
