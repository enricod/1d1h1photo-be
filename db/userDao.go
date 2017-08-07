package db

import (
	"github.com/jinzhu/gorm"
	"errors"
)


type User struct {
	gorm.Model
	Username string
	Email string
	EmailValid bool
}

type UserAppToken struct {
	gorm.Model
	UserId uint
	AppToken string
	Valid bool
}

func UserFindByEmail(  email string)  (User, error) {
	db := openDB();
	var user User
	db.First(&user, "email=?", email) // find product with code l1212
	if (user.ID > 0) {
		return user, nil
	} else {
		return user, errors.New("nessun utente con email ")
	}
}

func SalvaUser(user *User) {
	db.Create(user)
}

func SalvaAppToken(userId uint, appToken string) {
	db.Create(&UserAppToken{UserId: userId, AppToken: appToken})
}
