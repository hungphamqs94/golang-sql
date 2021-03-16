package singleton

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func GetInstance() *sql.DB {
	if db == nil {
		db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
		if err != nil {
			panic(err.Error())
		}
	}
	return db
}
