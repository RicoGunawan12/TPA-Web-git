package model

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   []byte `json:"password"`
	Phone      string `json:"phone"`
	RoleID     int    `json:"role_id" gorm:"foreign_key: role_id"`
	Subscribed bool   `json:"subscribed"`
	Location   string `json:"location"`
	Money      int    `json:"money"`
	Status     string `json:"status"`
	ShopID     int    `json:"shop_id"`
}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
