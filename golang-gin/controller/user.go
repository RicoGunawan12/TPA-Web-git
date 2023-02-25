package controller

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/RicoGunawan12/gin-gorm-api/config"
	"github.com/RicoGunawan12/gin-gorm-api/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(ctx *gin.Context) {
	users := []model.User{}
	config.DB.Where("role_id = 1").Find(&users)
	ctx.JSON(200, &users)
}

func GetUser(ctx *gin.Context) {
	users := []model.User{}
	config.DB.Where("id = ?", ctx.Param("id")).Find(&users)
	ctx.JSON(200, &users)
}

func GetUserPaginate(ctx *gin.Context) {
	users := []model.User{}
	config.DB.Where("role_id = 1").Order("id ASC").Find(&users)
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize := 8

	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	if endIndex > len(users) {
		endIndex = len(users)
	}

	if startIndex >= endIndex {
		ctx.String(200, "All data fetched")
		return
	}
	data := users[startIndex:endIndex]

	ctx.JSON(200, &data)
}

func CreateUser(ctx *gin.Context) {
	var user model.User
	ctx.ShouldBindJSON(&user)

	var emailCount int64 = 0
	config.DB.Model(model.User{}).Where("email = ?", user.Email).Count(&emailCount)

	if emailCount > 0 {
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		panic(err)
	}

	user.Password = password

	config.DB.Create(&user)
	ctx.JSON(200, &user)
}

func SignIn(ctx *gin.Context) {
	var attempt model.Cred
	ctx.ShouldBindJSON(&attempt)
	fmt.Println(attempt.Email)
	fmt.Println(attempt.Password)

	var user model.User
	config.DB.Model(model.User{}).Where("email = ?", attempt.Email).First(&user)

	if user.ID == 0 {
		ctx.String(200, "Email Not Found")
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(attempt.Password)); err != nil {
		ctx.String(200, "Incorrect Password")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.MyJWTClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.FirstName + " " + user.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(os.Getenv("secretKey")))

	if err != nil {
		ctx.JSON(400, "error")
	}

	ctx.SetCookie("jwt", ss, 3600, "/", "http://localhost:3000", false, true)
	// fmt.Print(ss)
	ctx.JSON(200, ss)
}

func Auth(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(200, user)
}

func SignOut(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, "Logout successful!")
}

func DeleteUser(ctx *gin.Context) {
	var user model.User
	config.DB.Where("id = ?", ctx.Param("id")).Delete(user)
}

func BanUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	var updatedUser *model.User
	config.DB.Where("id = ?", userID).First(&updatedUser)
	updatedUser.Status = "banned"
	config.DB.Save(&updatedUser)
	ctx.String(http.StatusOK, "User banned!")

}

func UnbanUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	var updatedUser *model.User
	config.DB.Where("id = ?", userID).First(&updatedUser)
	updatedUser.Status = "active"
	config.DB.Save(&updatedUser)
	ctx.String(http.StatusOK, "User banned!")

}

func SendEmail(ctx *gin.Context) {
	type ToSend struct {
		EmailToSend string `json:"email_to_send"`
		Header      string `json:"header"`
		Body        string `json:"body"`
	}
	var sendUser ToSend
	ctx.BindJSON(&sendUser)

	// Set up authentication information for SMTP server
	auth := smtp.PlainAuth("", "ricogunawan12323@gmail.com", "dubhmseksbteatbl", "smtp.gmail.com")

	// Set up message headers
	// headers := make(map[string]string)
	// headers["From"] = "ricogunawan12323@gmail.com"
	// headers["To"] = email
	// headers["Subject"] = subject

	// Build message body
	// body := []byte(sendUser.Body)

	headers := make(map[string]string)
	headers["From"] = "ricogunawan12323@gmail.com"
	headers["To"] = sendUser.EmailToSend
	headers["Subject"] = sendUser.Header

	message := ""
	for key, value := range headers {
		message += key + ": " + value + "\r\n"
	}

	message += "\r\n" + sendUser.Body

	// Send the email
	err := smtp.SendMail("smtp.gmail.com:587", auth, "ricogunawan12323@gmail.com", []string{sendUser.EmailToSend}, []byte(message))
	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, gin.H{
		// 	"error": "Failed to send email",
		// })
		// return
		fmt.Print(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Email sent successfully",
	})
}
