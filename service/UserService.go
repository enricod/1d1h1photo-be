package service

import (
	"errors"

	"1d1hphoto.com-be/convert"
	"1d1hphoto.com-be/db"
	"1d1hphoto.com-be/model"
)

func GetUser(user string) (model.User, error) {
	var rows = db.ExecuteSelect("select * from user where id=" + user)
	var result = convert.ParseUsers(rows)
	if len(result) > 0 {
		return result[0], nil
	}
	return model.User{}, errors.New("User non trovato")
}

func ListUser() []model.User {
	var rows = db.ExecuteSelect("select * from user")
	return convert.ParseUsers(rows)
}
