package todo

import "time"

type DataResponse struct {
	Id              uint      `json:"id"`
	Title           string    `json:"title"`
	ActivityGroupID int64     `json:"activity_group_id"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
