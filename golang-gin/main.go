package main

import (
	"github.com/RicoGunawan12/gin-gorm-api/config"
	"github.com/RicoGunawan12/gin-gorm-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods:     []string{"POST", "GET", "PUT", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")
	corsConfig.AddAllowMethods("GET")
	corsConfig.AddAllowMethods("POST")
	corsConfig.AddAllowMethods("PUT")
	corsConfig.AddAllowMethods("PATCH")

	// Register the middleware
	router.Use(cors.New(corsConfig))

	config.Connect()
	routes.UserRoute(router)
	routes.ProductRoute(router)
	routes.VoucherRoute(router)
	routes.PromotionRoute(router)
	routes.ShopRoute(router)
	routes.CartRoute(router)

	router.Run(":8080")
}

// func CORSAccess(router *gin.Engine) {

// }
