package model

import "gorm.io/gorm"

type ProductSubCategory struct {
	gorm.Model
	ProductSubCategory string `json:"product_sub_category"`
}
