package db

import (
	"github.com/enricod/1h1dphoto.com-be/model"
	"errors"
)


func UserFindByEmail(  email string)  (model.User, error) {
	db := openDB();
	var user model.User
	db.First(&user, "email=?", email) // find product with code l1212
	if (user.ID > 0) {
		return user, nil
	} else {
		return user, errors.New("nessun utente con email ")
	}
}


func ValidateUserAppToken( validationCode string, appToken string) bool  {
	db := openDB();
	var userAppToken model.UserAppToken
	db.First(&userAppToken, "validation_code = ? AND app_token=?", validationCode, appToken)
	if userAppToken.ID > 0 {
		userAppToken.Valid = true
		db.Save(&userAppToken)
		return true
	} else {
		return false
	}

}

func FindAppToken(appToken string)  *model.UserAppToken  {
	db := openDB();
	var userAppToken model.UserAppToken
	db.First(&userAppToken, " app_token=?", appToken)
	if userAppToken.ID > 0 {
		return &userAppToken
	} else {
		return nil
	}

}

func SalvaUser(user *model.User) {
	db.Create(user)
}

func SalvaAppToken(userId uint, appToken string, validationCode string) {
	var user model.User
	db.First(&user, userId)
	db.Create(&model.UserAppToken{User: user, AppToken: appToken, ValidationCode: validationCode})
}
