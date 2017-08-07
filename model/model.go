package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/enricod/1h1dphoto.com-be/db"
	"math/rand"
)

var MyKey = "1d1hphoto"

type Claims struct {
	Username string
	Time     int64
	jwt.StandardClaims
}

type Response struct {
	Status bool
	Msg    string
	Data   string
}




/**
 primo avvio del client
 */
type UserRegisterReq struct {
	Username string
	Email string
}

type UserCodeValidationReq struct {
	ValidationCode string
	UserToken string
}

type ResHead struct {
	Success bool `json:"success"`
}

type UserRegisterResBody struct {
	UserToken string `json:"userToken"`
	User db.User
}

type UserRegisterRes struct {
	Head ResHead `json:"head"`
	Body UserRegisterResBody `json:"body"`
}


const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
