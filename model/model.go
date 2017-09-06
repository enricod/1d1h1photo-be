package model

import (
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

// Confs configurazioni generali applicazione
var Confs AppConfs

// AppConfs configurazioni dell'app, lette da file di configurazione esterno
type AppConfs struct {
	Port         int
	ImgDir       string
	ImgUploadDir string
	DbUser       string
	DbPass       string
	DbDatabase   string
}

// Response probabilmente obsoleto
type Response struct {
	Status bool
	Msg    string
	Data   string
}

// UserRegisterReq request primo avvio dell'applicazione mobile
type UserRegisterReq struct {
	Username string
	Email    string
}

// UserCodeValidationReq req per validazione codice utente
type UserCodeValidationReq struct {
	ValidationCode string
	AppToken       string
}

// ResHead sezione head delle risposte JSON
type ResHead struct {
	Success bool `json:"success"`
}

// UserRegisterResBody body per req di registrazione utente
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

// IsClosed true se evento è terminato
func (e Event) IsClosed() bool {
	return e.End.Before(time.Now())
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

// Msg per traduzioni
type Msg struct {
	gorm.Model
	Code   string
	Locale string
	Value  string
	Desc   string
}

// User utente da database
type User struct {
	gorm.Model
	Username   string
	Email      string
	EmailValid bool
}

// UserAppToken contiene tokens associati ad utente
type UserAppToken struct {
	gorm.Model
	User           User `gorm:"ForeignKey:UserId"`
	UserId         uint
	AppToken       string
	ValidationCode string
	Valid          bool
}

type EventsSummaryResBody struct {
	FutureEvents []Event `json:"futureEvents"`
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

// GenerateRandomBytes generazione bytes casuali
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
