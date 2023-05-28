// this uploads the user data to the database

package handlers

import (
	"path/filepath"
	"strconv"
	"wellfyn/db"
	"wellfyn/models"

	"github.com/gin-gonic/gin"
)

func UploadData(c *gin.Context) {
	user, _ := c.Keys["user"].(models.Userauth)
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	var imagesString string    // to be stored in the database
	var documentsString string // to be stored in the database

	form, err := c.MultipartForm() // to accept multiple files

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error in uploading data",
		})
		return
	}

	images := form.File["images"]
	documents := form.File["documents"]

	// saving the files to the server
	for _, image := range images {
		filename := filepath.Base(image.Filename)
		filepathString := "usermedia/" + strconv.FormatUint(uint64(user.ID), 10) + "/" + filename
		if err := c.SaveUploadedFile(image, filepathString); err != nil {
			c.JSON(400, gin.H{
				"message": "Error in uploading data",
			})
			return
		} else {
			imagesString += filepathString + ","
		}
	}

	// saving the documents on the server
	for _, document := range documents {
		filename := filepath.Base(document.Filename)
		filePathString := "usermedia/" + strconv.FormatUint(uint64(user.ID), 10) + "/" + filename
		if err := c.SaveUploadedFile(document, filePathString); err != nil {
			c.JSON(400, gin.H{
				"message": "Error in uploading data",
			})
			return
		} else {
			documentsString += filePathString + ","
		}
	}
	// uploading the data to the database and handling the error
	result := db.DB.Table("user_data").Where("email = ?", user.Email).Save(&models.UserData{Name: name, Phone: phone, Images: imagesString, Documents: documentsString, Email: user.Email})

	// handling the error

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error in uploading data to sql",
		})
		return
	}

	// final response
	c.JSON(200, gin.H{
		"message": "Data uploaded successfully",
		"user":    user,
	},
	)

}
