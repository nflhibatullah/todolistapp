package entity

import "gorm.io/gorm"

type Activity struct {
		gorm.Model
		Title string `json:"title"`
		Email string `json:"email"`
	}


