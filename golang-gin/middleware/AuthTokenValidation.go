package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/RicoGunawan12/gin-gorm-api/config"
	"github.com/RicoGunawan12/gin-gorm-api/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func ValidateToken(ctx *gin.Context) {
	type JWT struct {
		gorm.Model
		Token string `json:"token"`
	}

	var token JWT
	ctx.ShouldBindJSON(&token)

	if token.Token == "" {
		ctx.String(200, "Couldn't Get Cookie")
		return
	}

	tokenString := token.Token
	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, success := token.Method.(*jwt.SigningMethodHMAC); !success {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("secretKey")), nil
	})

	if err != nil {
		ctx.String(200, "Token Parsing Failed")
		return
	}

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {

			ctx.String(200, "Cookie Expired")
			return

		}
		var user model.User
		config.DB.First(&user, "ID = ?", claims["id"])

		if user.ID == 0 {
			ctx.String(200, "Email Not Found")
			return
		}

		ctx.Set("user", user)

		ctx.Next()

	} else {
		ctx.String(200, "Server Error")
	}
}
