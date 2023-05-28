package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // this is the global variable which we will use to interact with the database

func Init() error {
	Connect() // connect to the database

	// create a new Gorm instance with the connected database
	DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: MysqlDb,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Gorm initialized")
	return nil
}
