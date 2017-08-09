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
	User User `gorm:"ForeignKey:UserId"`
	UserId uint
	AppToken string
	ValidationCode string
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


func ValidateUserAppToken( validationCode string, appToken string) bool  {
	db := openDB();
	var userAppToken UserAppToken
	db.First(&userAppToken, "validation_code = ? AND app_token=?", validationCode, appToken)
	if userAppToken.ID > 0 {
		userAppToken.Valid = true
		db.Save(&userAppToken)
		return true
	} else {
		return false
	}

}

func FindAppToken(appToken string)  *UserAppToken  {
	db := openDB();
	var userAppToken UserAppToken
	db.First(&userAppToken, " app_token=?", appToken)
	if userAppToken.ID > 0 {
		return &userAppToken
	} else {
		return nil
	}

}

func SalvaUser(user *User) {
	db.Create(user)
}

func SalvaAppToken(userId uint, appToken string, validationCode string) {
	var user User
	db.First(&user, userId)
	db.Create(&UserAppToken{User: user, AppToken: appToken, ValidationCode: validationCode})
}
