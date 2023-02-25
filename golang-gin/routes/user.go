package routes

import (
	"github.com/RicoGunawan12/gin-gorm-api/controller"
	"github.com/RicoGunawan12/gin-gorm-api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controller.GetUsers)
	router.GET("/users/paginate", controller.GetUserPaginate)
	// router.GET("/api/users/:id", controller.GetUser)
	// router.POST("/api/users", controller.CreateUser)
	// router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/user/ban/:id", controller.BanUser)
	router.PUT("/user/unban/:id", controller.UnbanUser)
	router.POST("/sign-up", controller.CreateUser)
	router.POST("/sign-in", controller.SignIn)
	router.GET("/sign-out", controller.SignOut)
	router.POST("/auth", middleware.ValidateToken, controller.Auth)
	router.POST("/users/send-email", controller.SendEmail)
}
