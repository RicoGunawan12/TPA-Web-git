package routes

import (
	"github.com/RicoGunawan12/gin-gorm-api/controller"
	"github.com/gin-gonic/gin"
)

func PromotionRoute(router *gin.Engine) {
	router.POST("/add-promotion", controller.AddPromotion)
	router.GET("/promotions", controller.GetPromotions)
	router.DELETE("/promotion/:id", controller.DeletePromotion)
}
