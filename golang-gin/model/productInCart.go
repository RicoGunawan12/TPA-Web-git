package model

import (
	"gorm.io/gorm"
)

type ProductInCart struct {
	gorm.Model
	ProductID int `json:"product_id"`
	CartID    int `json:"cart_id"`
}
