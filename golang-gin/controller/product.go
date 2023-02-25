package controller

import (
	"fmt"
	"strconv"

	"github.com/RicoGunawan12/gin-gorm-api/config"
	"github.com/RicoGunawan12/gin-gorm-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProducts(ctx *gin.Context) {
	products := []model.Product{}
	config.DB.Find(&products)
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize := 20

	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	if endIndex > len(products) {
		endIndex = len(products)
	}

	if startIndex > endIndex {
		ctx.String(200, "All data fetched")
		return
	}
	data := products[startIndex:endIndex]

	ctx.JSON(200, &data)
}

func GetSellerProducts(ctx *gin.Context) {
	products := []model.Product{}
	config.DB.Where("shop_id = ?", ctx.Param("id")).Find(&products)
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize := 50
	fmt.Print(products)

	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	if endIndex > len(products) {
		endIndex = len(products)
	}

	if startIndex > endIndex {
		ctx.String(200, "All data fetched")
		return
	}
	data := products[startIndex:endIndex]

	ctx.JSON(200, &data)
}

func AddProduct(ctx *gin.Context) {
	var product model.Product
	ctx.BindJSON(&product)
	config.DB.Create(&product)
	ctx.String(200, "Insert Success!")
}

func GetProduct(ctx *gin.Context) {
	type JoinedProduct struct {
		gorm.Model
		ProductCategoryID          int     `json:"product_category_id"`
		ProductSubCategoryID       int     `json:"product_sub_category_id"`
		ProductSubCategoryDetailID int     `json:"product_sub_category_detail_id"`
		ProductID                  int     `json:"product_id"`
		BrandID                    int     `json:"brand_id"`
		ProductName                string  `json:"product_name"`
		ProductPrice               int     `json:"product_price"`
		Rating                     float64 `json:"rating"`
		SoldQuantity               int     `json:"sold_quantity"`
		Stock                      int     `json:"stock"`
		ShopID                     int     `json:"shop_id"`
		NumberOfReview             int     `json:"number_of_review"`
		Description                string  `json:"description"`
		PowerUsage                 int     `json:"power_usage"`
		ImageLink                  string  `json:"image_link"`
		PowerSupply                int     `json:"power_supply"`
		ProductSubCategoryDetail   string  `json:"product_sub_category_detail"`
		ProductSubCategory         string  `json:"product_sub_category"`
		ProductCategory            string  `json:"product_category"`
		Brand                      string  `json:"brand"`
		ShopName                   string  `json:"shop_name"`
		ShopPicture                string  `json:"shop_picture"`
	}

	var product JoinedProduct

	config.DB.Model(&model.Product{}).
		Joins("left join product_categories on products.product_category_id = product_categories.id").
		Joins("left join product_sub_categories on products.product_sub_category_id = product_sub_categories.id").
		Joins("left join product_sub_category_details on products.product_sub_category_detail_id = product_sub_category_details.id").
		Joins("left join brands on products.brand_id = brands.id").
		Joins("left join shops on products.shop_id = shops.id").
		Select("products.*, product_categories.product_category, product_sub_categories.product_sub_category, product_sub_category_details.product_sub_category_detail, brands.brand, shops.shop_name, shops.shop_picture").Where("products.id = ?", ctx.Param("id")).Find(&product)

	// config.DB..Find(&product)
	ctx.JSON(200, &product)
}

func GetTop3(ctx *gin.Context) {
	var top3 []model.Product
	config.DB.Order("sold_quantity desc").Limit(3).Find(&top3)
	ctx.JSON(200, &top3)
}
