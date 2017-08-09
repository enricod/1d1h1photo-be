package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/enricod/1h1dphoto.com-be/db"
	"math/rand"
	"time"
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
 * primo avvio del client
 */
type UserRegisterReq struct {
	Username string
	Email string
}

type UserCodeValidationReq struct {
	ValidationCode string
	AppToken string
}


type ResHead struct {
	Success bool `json:"success"`
}

type UserRegisterResBody struct {
	AppToken string `json:"appToken"`
	User db.User
}

type UserRegisterRes struct {
	Head ResHead `json:"head"`
	Body UserRegisterResBody `json:"body"`
}

type EventsSummaryResBody struct {
	NextEvent db.Event `json:"nextEvent"`
	ClosedEvents []db.Event `json:"closedEvents"`
}

type EventsListRes struct {
	Head ResHead `json:"head"`
	Body EventsSummaryResBody `json:"body"`
}

func GenerateRandomBytes(n int) ([]byte, error) {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}
