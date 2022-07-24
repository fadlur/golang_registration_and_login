package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
	Description string `json:"description"`
}

func (Role) TableName() string {
	return "role"
}