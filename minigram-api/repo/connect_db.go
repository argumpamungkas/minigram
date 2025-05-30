package repo

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MYSQL
// var (
// 	username = "root"
// 	//password = ""
// 	dbName = "minigram-db"
// 	host   = "127.0.0.2:3306"
// 	db     *gorm.DB
// 	err    error
// )

// POSTGRES
var (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	pwd    = "argum12"
	dbName = "minigram-db"
	db     *gorm.DB
	err    error
)

func StartDB() {
	// MYSQL
	// dsn := fmt.Sprintf("%s:@tcp(%s)/%s?parseTime=true&loc=Local", username, host, dbName)

	// POSTGRES
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, pwd, dbName, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
