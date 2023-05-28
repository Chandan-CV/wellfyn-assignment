package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomePageHandler handles the home page route

func HomePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "Hello World! this is the home page, navigate to /signup to signup and /login to login",
	})

}
