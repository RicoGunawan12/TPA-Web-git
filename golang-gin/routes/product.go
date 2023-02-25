package routes

import (
	"github.com/RicoGunawan12/gin-gorm-api/controller"
	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.Engine) {
	router.GET("/products", controller.GetProducts)
	router.GET("/get-product/:id", controller.GetProduct)
	router.GET("/seller-products/:id", controller.GetSellerProducts)
	router.POST("/add-product", controller.AddProduct)
	router.GET("/get-top-3", controller.GetTop3)
}
