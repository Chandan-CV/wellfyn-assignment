package handlers

import (
	"fmt"
	"wellfyn/db"
	"wellfyn/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(c *gin.Context) {
	//reading the post request body
	var body models.UserSignup
	c.Bind(&body)
	fmt.Println(body)

	// check if user exists
	var usercount int64
	db.DB.Table("userauths").Where("email = ?", body.Email).Count(&usercount)
	if usercount > 0 {
		c.JSON(400, gin.H{
			"message": "User already exists",
		})
		return
	}

	// hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error, could not hash password",
			"error":   err,
		})
		return
	}

	// if user does not exist, create user
	db.DB.Table("userauths").Create(&models.Userauth{Email: body.Email, Password: string(hashedPassword)})
	c.JSON(200, gin.H{
		"message": "User created successfully",
	})
	return

}
