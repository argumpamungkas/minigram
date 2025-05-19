package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

const (
	username = "root"
	password = ""
	dbName   = "minigram-db"
	host     = "127.0.0.2:3306"
)

func StartDB() {
	sqlInfo := fmt.Sprintf("%s:@tcp(%s)/%s", username, host, dbName)

	db, err = sql.Open("mysql", sqlInfo)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success connect to DB")
}

func GetDb() *sql.DB {
	return db
}
