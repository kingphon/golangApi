package responsemodel

type (
	PermissionCreate struct {
		ID string `json:"_id"`
	}

	PermissionUpdate struct {
		ID string `json:"_id"`
	}

	PermissionAll struct {
		Data []string `json:"data"`
	}
)
