package model

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	ShopName            string  `json:"shop_name"`
	ShopEmail           string  `json:"shop_email"`
	ShopPassword        string  `json:"shop_password"`
	Status              string  `json:"status"`
	DeliveryStatistic   float64 `json:"delivery_statistic"`
	Banner              string  `json:"banner"`
	ShopPicture         string  `json:"shop_picture"`
	ProductAccuracy     float64 `json:"product_accuracy"`
	ServiceSatisfaction float64 `json:"service_satisfaction"`
	NumberOfSales       int     `json:"number_of_sales"`
	Description         string  `json:"description"`
}
