package routes

import (
	"github.com/RicoGunawan12/gin-gorm-api/controller"
	"github.com/gin-gonic/gin"
)

func ShopRoute(router *gin.Engine) {
	router.POST("/insert-shop", controller.AddShop)
	router.GET("/shops", controller.GetShops)
	router.GET("/shop/:id", controller.GetShop)
	router.PUT("/shop/ban/:id", controller.BanShop)
	router.PUT("/shop/unban/:id", controller.UnbanShop)
}
