package main

import (
	"wellfyn/db"
	"wellfyn/handlers"
	"wellfyn/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	// initializing database and Migrating it
	db.Init()
	db.MigrateDB()

	// starting the server and setting up routes
	r := gin.Default()
	r.Static("/usermedia", "./usermedia")
	r.GET("/", handlers.HomePageHandler)
	r.POST("/signup", handlers.SignUpHandler)
	r.POST("/login", handlers.Login)
	r.GET("/validate", middlewares.ReqAuth, handlers.Validate)

	r.MaxMultipartMemory = 8 << 20 // 8 MiB is the max size of the file accepted

	r.POST("/upload", middlewares.ReqAuth, handlers.UploadData)

	r.Run()
}
