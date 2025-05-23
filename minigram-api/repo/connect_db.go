package repo

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	username = "root"
	//password = ""
	dbName = "minigram-db"
	host   = "127.0.0.2:3306"
	db     *gorm.DB
	err    error
)

func StartDB() {
	dsn := fmt.Sprintf("%s:@tcp(%s)/%s?parseTime=true", username, host, dbName)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// not use gorm
	// db, err = sql.Open("mysql", sqlInfo)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err)
	}

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Success connect to DB")
}

func GetDb() *gorm.DB {
	return db
}
