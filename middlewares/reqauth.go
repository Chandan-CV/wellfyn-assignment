// this is an important middleware that checks if the user is authenticated or not

package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"wellfyn/db"
	"wellfyn/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ReqAuth(c *gin.Context) {

	// get the token from the header
	cookie, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token not found, please login",
		})
		return
	}

	// check if the token is valid

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("SECRET")), nil
		// SECRET is the secret key for the password
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token, please login again",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)

		// check if the token is expired
		exp := claims["exp"].(float64)

		if float64(time.Now().Unix()) > (exp) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token expired",
			})

		}

		// find user by id
		var user models.Userauth
		db.DB.Table("userauths").First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})

			return
		}
		fmt.Println(user)
		c.Set("user", user)

		c.Next()

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token, please login again",
		})
		return
	}
}
