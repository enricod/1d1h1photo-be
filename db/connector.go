package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_HOST = "tcp(127.0.0.1:8889)"
	DB_NAME = "test"
	DB_USER = "root"
	DB_PASS = "root"
)

func ExecuteSelect(query string) *sql.Rows {
	db := openDB()
	rows, err := db.Query(query)
	CheckErr(err)
	db.Close()
	return rows
}

func Insert(query string, s ...string) int64 {
	db := openDB()
	stmt, err := db.Prepare(query)
	CheckErr(err)
	res, err := stmt.Exec("w", "e")
	CheckErr(err)
	id, err := res.LastInsertId()
	CheckErr(err)
	db.Close()
	return id
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@"+DB_HOST+"/"+DB_NAME+"?charset=utf8")
	CheckErr(err)
	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
