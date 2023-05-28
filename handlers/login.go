// this is the login page handler

package handlers

import (
	"net/http"
	"os"
	"time"
	"wellfyn/db"
	"wellfyn/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	// get the email and pass of the body

	var body models.UserSignup

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//look up requested user
	var user models.Userauth

	result := db.DB.Table("userauths").Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	//compare the sent password with the stored hash

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	//Generate a token

	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte(os.Getenv("SECRET"))
	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	)
	s, err = t.SignedString(key)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	//send the token to the user as a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", s, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "user logged in! and token sent as a cookie"})

}
