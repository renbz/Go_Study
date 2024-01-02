package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:ren372930@tcp(82.157.50.241:3306)/bookstore")
	if err != nil {
		panic(err.Error())
	}
}
