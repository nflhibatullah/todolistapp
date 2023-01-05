package entity

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ActivityGroupID int64  `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active" gorm:"default:true"`
	Priority        string `json:"priority" gorm:"default:'very-high'"`
}
