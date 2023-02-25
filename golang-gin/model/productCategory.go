package model

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	ProductCategory string `json:"product_category"`
}
