package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var user string = "<your username>"
var password string = "<your db password>"
var host string = "<your host>"
var dbName string = "<your DB Name>"

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", (user + ":" + password + "@tcp(" + host + ")/" + dbName + ")"))

	if err != nil {
		panic(err)
	}

	return db
}
