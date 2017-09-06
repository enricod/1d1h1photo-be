package db

import (
	//_ "github.com/go-sql-driver/mysql"
	"github.com/enricod/1h1dphoto.com-be/model"
	"github.com/jinzhu/gorm"
	// import per driver mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func openDB() *gorm.DB {

	if db == nil {
		driverName := "mysql"
		aDb, err := gorm.Open(driverName, model.Confs.DbUser+":"+model.Confs.DbPass+"@/"+
			model.Confs.DbDatabase+
			"?charset=utf8&parseTime=True&loc=Local")

		db = aDb
		db.LogMode(true)

		checkErr(err)

		// Migrate the schema
		db.AutoMigrate(&model.User{}, &model.UserAppToken{}, &model.Event{},
			&model.EventSubmission{}, &model.EventSubmissionAction{})

	}
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
