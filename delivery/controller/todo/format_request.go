package todo

type CreateToDoRequest struct {
	Title           string `json:"title"`
	ActivityGroupID int64  `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
}

type UpdateToDoRequest struct {
	Title           string `json:"title"`
	ActivityGroupID int64  `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}
