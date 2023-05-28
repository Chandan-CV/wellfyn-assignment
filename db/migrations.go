// files to migrage all the structs to the database
package db

import (
	"fmt"
	"wellfyn/models"
)

func MigrateDB() {
	fmt.Println("Migrating DB")
	DB.AutoMigrate(&models.Userauth{})
	DB.AutoMigrate(&models.UserData{})
	fmt.Println("DB Migrated")
}
