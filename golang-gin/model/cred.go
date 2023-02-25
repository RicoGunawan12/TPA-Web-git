package model

import "gorm.io/gorm"

type Cred struct {
	gorm.Model
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
