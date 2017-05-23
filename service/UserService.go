package service

import (
	"errors"

	"fmt"

	"1d1hphoto.com-be/convert"
	"1d1hphoto.com-be/db"
	"1d1hphoto.com-be/model"
)

func GetUser(user string) (model.User, error) {
	query := fmt.Sprintf("select * from user where id=\"%s\"", user)
	fmt.Println("esegui query: ", query)
	var rows = db.ExecuteSelect(query)
	var result = convert.ParseUsers(rows)
	if len(result) > 0 {
		return result[0], nil
	}
	return model.User{}, errors.New("User non trovato")
}

func GetUserByUsername(user string) (model.User, error) {
	query := fmt.Sprintf("select * from user where username=\"%s\"", user)
	fmt.Println("esegui query: ", query)
	var rows = db.ExecuteSelect(query)
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
