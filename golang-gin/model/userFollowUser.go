package model

import (
	"gorm.io/gorm"
)

type UserFollowUser struct {
	gorm.Model
	TargetID int `json:"target_id"`
	UserID   int `json:"user_id"`
}
