package model

import (
	"gorm.io/gorm"
)

type UserFollowWishList struct {
	gorm.Model
	WishlistID int `json:"wishlist_id"`
	UserID     int `json:"user_id"`
}
