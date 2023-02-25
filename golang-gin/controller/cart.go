package controller

import (
	"github.com/RicoGunawan12/gin-gorm-api/config"
	"github.com/RicoGunawan12/gin-gorm-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddToCart(ctx *gin.Context) {
	var cart model.Cart
	ctx.ShouldBindJSON(&cart)
	config.DB.Create(&cart)
	ctx.String(200, "Add Success!")
}

func GetCart(ctx *gin.Context) {
	type Cart struct {
		gorm.Model
		UserID         int     `json:"user_id"`
		ProductID      int     `json:"product_id"`
		Quantity       int     `json:"quantity"`
		ProductName    string  `json:"product_name"`
		ProductPrice   int     `json:"product_price"`
		Rating         float64 `json:"rating"`
		SoldQuantity   int     `json:"sold_quantity"`
		Stock          int     `json:"stock"`
		ShopID         int     `json:"shop_id"`
		NumberOfReview int     `json:"number_of_review"`
		Description    string  `json:"description"`
		PowerUsage     int     `json:"power_usage"`
		ImageLink      string  `json:"image_link"`
		PowerSupply    int     `json:"power_supply"`
	}
	var cart []Cart

	config.DB.
		Joins("join products on products.ID = carts.product_id").
		Select("carts.*, products.*").
		Where("user_id = ?", ctx.Param("id")).Find(&cart)
	ctx.JSON(200, &cart)
}

func DeleteProductFromCart(ctx *gin.Context) {
	var temp model.Cart
	ctx.BindJSON(&temp)
	var temp2 model.Cart
	config.DB.Where("user_id = ?", temp.UserID).Where("product_id = ?", temp.ID).Find(&temp2)
	config.DB.Delete(&temp2)
}
