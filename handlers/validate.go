// A simple handler to validate the token and check if the user is logged in or not.
package handlers

import "github.com/gin-gonic/gin"

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, gin.H{
		"user": user,
	})
}
