package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleID int    `json:"role_id" gorm:"primary_key"`
	Role   string `json:"role"`
}
