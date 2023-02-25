package controller

import (
	"github.com/RicoGunawan12/gin-gorm-api/config"
	"github.com/RicoGunawan12/gin-gorm-api/model"
	"github.com/gin-gonic/gin"
)

func AddPromotion(ctx *gin.Context) {
	var promotion model.Promotion
	ctx.ShouldBindJSON(&promotion)
	config.DB.Create(&promotion)
	ctx.String(200, "Add Success!")
}

func GetPromotions(ctx *gin.Context) {
	promotions := []model.Promotion{}
	config.DB.Find(&promotions)
	ctx.JSON(200, &promotions)
}

func DeletePromotion(ctx *gin.Context) {
	var promotion model.Promotion
	config.DB.Where("id = ?", ctx.Param("id")).Delete(&promotion)
}
