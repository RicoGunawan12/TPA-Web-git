package model

import "gorm.io/gorm"

type ProductSubCategoryDetail struct {
	gorm.Model
	ProductSubCategoryDetail string `json:"product_sub_category_detail"`
}
