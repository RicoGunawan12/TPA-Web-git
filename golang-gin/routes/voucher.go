package routes

import (
	"github.com/RicoGunawan12/gin-gorm-api/controller"
	"github.com/gin-gonic/gin"
)

func VoucherRoute(router *gin.Engine) {
	router.POST("/add-voucher", controller.AddVoucher)
}
