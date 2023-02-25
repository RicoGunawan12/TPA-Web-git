package routes

import (
	"github.com/RicoGunawan12/gin-gorm-api/controller"
	"github.com/gin-gonic/gin"
)

func CartRoute(route *gin.Engine) {
	route.POST("/add-to-cart", controller.AddToCart)
	route.GET("/get-cart/:id", controller.GetCart)
	route.POST("/delete-from-cart", controller.DeleteProductFromCart)
}
