package controller

import (
	"github.com/RicoGunawan12/gin-gorm-api/config"
	"github.com/RicoGunawan12/gin-gorm-api/model"
	"github.com/gin-gonic/gin"
)

func AddShop(ctx *gin.Context) {
	var shop model.Shop
	ctx.BindJSON(&shop)

	var shopCount model.Shop
	config.DB.Where("shop_name = ?", shop.ShopName).First(&shopCount)

	if shopCount.ID != 0 {
		ctx.String(200, "This name is taken!")
		return
	} else {
		config.DB.Create(&shop)
		ctx.JSON(200, &shop)
	}

}

func GetShops(ctx *gin.Context) {
	shops := []model.Shop{}
	config.DB.Order("id asc").Find(&shops)
	ctx.JSON(200, &shops)
}

func GetShop(ctx *gin.Context) {
	shops := model.Shop{}
	config.DB.Where("id = ?", ctx.Param("id")).Find(&shops)
	ctx.JSON(200, &shops)
}

func BanShop(ctx *gin.Context) {
	shopID := ctx.Param("id")
	var updatedShop *model.Shop
	config.DB.Where("id = ?", shopID).First(&updatedShop)
	updatedShop.Status = "Banned"
	config.DB.Save(&updatedShop)

	var updatedUser *model.User
	config.DB.Where("shop_id = ?", shopID).First(&updatedUser)
	updatedUser.Status = "Banned"
	config.DB.Save(&updatedUser)

}

func UnbanShop(ctx *gin.Context) {
	shopID := ctx.Param("id")
	var updatedShop *model.Shop
	config.DB.Where("id = ?", shopID).First(&updatedShop)
	updatedShop.Status = "Active"
	config.DB.Save(&updatedShop)

	var updatedUser *model.User
	config.DB.Where("shop_id = ?", shopID).First(&updatedUser)
	updatedUser.Status = "Active"
	config.DB.Save(&updatedUser)
}
