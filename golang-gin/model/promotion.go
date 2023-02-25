package model

import "gorm.io/gorm"

type Promotion struct {
	gorm.Model
	PromotionImage       string `json:"promotion_image"`
	PromotionDescription string `json:"promotion_description"`
}
