// mysqldb connection
package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var MysqlDb *sql.DB // this is the global variable which we will use to interact with the database to create the gorm instance
var err error

func Connect() error {
	// config file
	config := mysql.Config{
		User:                 "root",
		Passwd:               "Development16",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "wellfyn",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	// connect to the database
	MysqlDb, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return err
	}
	// check if the connection is established
	pingErr := MysqlDb.Ping()
	if pingErr != nil {
		return (pingErr)
	}
	// if the connection is established, print the message
	fmt.Println("Successfully connected to database")
	return nil
}
