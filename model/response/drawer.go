package responsemodel

type (
	DrawerCreate struct {
		ID string `json:"_id"`
	}

	DrawerUpdate struct {
		ID string `json:"_id"`
	}

	DrawerUpdateStatus struct {
		ID string `json:"_id"`
	}

	DrawerAll struct {
		Limit int64         `json:"limit"`
		Total int64         `json:"total"`
		Data  []DrawerTable `json:"data"`
	}

	DrawerTable struct {
		ID        string       `json:"_id"`
		Cabinet   CabinetTable `json:"cabinet"`
		Name      string       `json:"name"`
		Active    string       `json:"active"`
		CreatedAt string       `json:"createdAt"`
	}
)
