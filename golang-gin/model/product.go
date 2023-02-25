package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductCategoryID          int     `json:"product_category_id"`
	ProductSubCategoryID       int     `json:"product_sub_category_id"`
	ProductSubCategoryDetailID int     `json:"product_sub_category_detail_id"`
	ProductID                  int     `json:"product_id"`
	BrandID                    int     `json:"brand_id"`
	ProductName                string  `json:"product_name"`
	ProductPrice               int     `json:"product_price"`
	Rating                     float64 `json:"rating"`
	SoldQuantity               int     `json:"sold_quantity"`
	Stock                      int     `json:"stock"`
	ShopID                     int     `json:"shop_id"`
	NumberOfReview             int     `json:"number_of_review"`
	Description                string  `json:"description"`
	PowerUsage                 int     `json:"power_usage"`
	ImageLink                  string  `json:"image_link"`
	PowerSupply                int     `json:"power_supply"`
}
