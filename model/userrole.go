package model

import "gorm.io/gorm"

type UserRole struct {	
	UserId int64 `json:"user_id"`
	RoleId int64 `json:"role_id"`
	gorm.Model
}

func (UserRole) TableName() string {
	return "user_role"
}