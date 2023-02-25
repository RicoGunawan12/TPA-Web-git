package config

import (
	"github.com/RicoGunawan12/gin-gorm-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=prk dbname=newegg port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Role{})
	db.AutoMigrate(&model.Cart{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.UserFollowUser{})
	db.AutoMigrate(&model.UserFollowWishList{})
	db.AutoMigrate(&model.ProductInCart{})
	db.AutoMigrate(&model.Voucher{})
	db.AutoMigrate(&model.Promotion{})
	db.AutoMigrate(&model.Shop{})
	db.AutoMigrate(&model.ProductCategory{})
	db.AutoMigrate(&model.ProductSubCategory{})
	db.AutoMigrate(&model.ProductSubCategoryDetail{})
	db.AutoMigrate(&model.Brand{})

	DB = db
}
