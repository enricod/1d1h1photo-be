package db

import (
	//_ "github.com/go-sql-driver/mysql"
	"github.com/enricod/1h1dphoto.com-be/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "onehouronedayphoto"
	DB_USER = "root"
	DB_PASS = "root"
)

var db *gorm.DB

func openDB() *gorm.DB {

	if db == nil {
		driverName := "mysql"
		aDb, err := gorm.Open(driverName, DB_USER+":"+DB_PASS+"@/"+DB_NAME+"?charset=utf8&parseTime=True&loc=Local")

		db = aDb
		db.LogMode(true)
		/*
			db, err := sql.Open(driverName, DB_USER + ":" + DB_PASS + "@" + DB_HOST +
					 "/" + DB_NAME + "?charset=utf8")
		*/

		CheckErr(err)

		//defer db.Close()

		// Migrate the schema
		db.AutoMigrate(&model.User{}, &model.UserAppToken{}, &model.Event{},
			&model.EventSubmission{}, &model.EventSubmissionAction{})

	}
	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
