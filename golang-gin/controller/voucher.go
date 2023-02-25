package controller

import (
	"github.com/RicoGunawan12/gin-gorm-api/config"
	"github.com/RicoGunawan12/gin-gorm-api/model"
	"github.com/gin-gonic/gin"
)

func AddVoucher(ctx *gin.Context) {
	var voucher model.Voucher
	ctx.ShouldBindJSON(&voucher)
	config.DB.Create(&voucher)
	ctx.String(200, "Add Success!")
}
