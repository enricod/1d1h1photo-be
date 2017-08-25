package model

import (
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

var MyKey = "1d1hphoto"

type Claims struct {
	Username string
	Time     int64
	jwt.StandardClaims
}

type AppConfs struct {
	Port         int
	ImgDir       string
	ImgUploadDir string
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
	Email    string
}

type UserCodeValidationReq struct {
	ValidationCode string
	AppToken       string
}

type ResHead struct {
	Success bool `json:"success"`
}

type UserRegisterResBody struct {
	AppToken string `json:"appToken"`
	User     User
}

type UserRegisterRes struct {
	Head ResHead             `json:"head"`
	Body UserRegisterResBody `json:"body"`
}

// Event oggetto dal database per descrizione evento
type Event struct {
	gorm.Model
	Start       time.Time
	End         time.Time
	Name        string
	Submissions []EventSubmission
}

// EventSubmission oggetto dal DB per descrizione immagine associata a evento
type EventSubmission struct {
	gorm.Model
	Event          Event `json:"-"`
	EventId        uint
	User           User `json:"-"`
	UserId         uint
	ImageName      string
	ImageUid       string
	ThumbUrl       string
	ImageUrl       string
	LikesNr        uint
	Score          float32
	Latitude       float32
	Longitude      float32
	SubmissionDate time.Time
}

// User utente da database
type User struct {
	gorm.Model
	Username   string
	Email      string
	EmailValid bool
}

type UserAppToken struct {
	gorm.Model
	User           User `gorm:"ForeignKey:UserId"`
	UserId         uint
	AppToken       string
	ValidationCode string
	Valid          bool
}

func (e Event) IsClosed() bool {
	return e.End.Before(time.Now())
}

type EventsSummaryResBody struct {
	NextEvent    Event   `json:"nextEvent"`
	ClosedEvents []Event `json:"closedEvents"`
}

type EventResBody struct {
	Head ResHead `json:"head"`
	Body Event   `json:"body"`
}

type EventsListRes struct {
	Head ResHead              `json:"head"`
	Body EventsSummaryResBody `json:"body"`
}

type EventSubmissionAction struct {
	gorm.Model
	User              User `gorm:"ForeignKey:UserId"`
	UserId            uint
	Time              time.Time
	EventSubmission   EventSubmission `gorm:"ForeignKey:UserId"`
	EventSubmissionId uint
	Type              string
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
