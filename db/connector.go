package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_HOST = "tcp(127.0.0.1:8889)"
	DB_NAME = "test"
	DB_USER = "root"
	DB_PASS = "root"
)

func Test() {
	rows, err := openDB().Query("SELECT * FROM user")
	CheckErr(err)
	for rows.Next() {
		var id int
		var nome string
		var cognome string
		err = rows.Scan(&id, &nome, &cognome)
		CheckErr(err)
		fmt.Println(id)
		fmt.Println(nome)
		fmt.Println(cognome)
	}
}

func ExecuteSelect(query string) *sql.Rows {
	rows, err := openDB().Query(query)
	CheckErr(err)
	return rows
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
