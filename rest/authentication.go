package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/enricod/1h1dphoto.com-be/db"
	"github.com/enricod/1h1dphoto.com-be/model"
	"github.com/gorilla/mux"
	"github.com/otium/queue"
	"gopkg.in/gomail.v2"
)

func version() string {
	return "0.2"
}

// Tokens è una mappa contenente token => utente
var Tokens = make(map[string]model.User)

type sender interface {
	send()
}

// EmailInfo struttura per informazione su utente
type EmailInfo struct {
	To   string
	Code string
}

func (ei EmailInfo) send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "1h1dphoto@gmail.com")
	m.SetHeader("To", ei.To)
	m.SetHeader("Subject", "validation code")
	m.SetBody("text/html", "Your validation code <b> "+ei.Code+"</b>")

	d := gomail.NewDialer("smtp.gmail.com", 465, "1h1dphoto@gmail.com", "1h1dphotos")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Errore in invio email %v", err)
		// panic(err)
	}
}

// Status chiamata per avere info sul backend
func Status(res http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(res).Encode(model.Response{Data: "Ok"})
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	} else {
		res.WriteHeader(http.StatusOK)
	}
}

// UserRegister riceve in post il tipo UserRegisterReq. Se utente non esiste già sul database, lo creo.
// genero codice alfanumerico che poi sarà spedito via email.
// Restituisce (in JSON) un istanza di UserRegisterRes
func UserRegister(res http.ResponseWriter, req *http.Request) {

	var userRegisterReq model.UserRegisterReq
	if req.Body == nil {
		http.Error(res, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&userRegisterReq)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	//fmt.Println(userRegisterReq.Username)

	// crea record per sessione
	validationCode, _ := model.GenerateRandomString(5)
	userToken, _ := model.GenerateRandomString(32)
	head := model.ResHead{Success: true}
	body := model.UserRegisterResBody{AppToken: userToken}
	userRegisterRes := model.UserRegisterRes{Head: head, Body: body}

	fmt.Println("validationCode: ", validationCode)

	// crea utente nel database se necessario
	if user, err := db.UserFindByEmail(userRegisterReq.Email); err != nil {
		user := model.User{Username: userRegisterReq.Username, Email: userRegisterReq.Email}
		db.SalvaUser(&user)
		// crea record in USER_APP_TOKEN
		db.SalvaAppToken(user.ID, userToken, validationCode)
	} else {
		// crea record in USER_APP_TOKEN
		db.SalvaAppToken(user.ID, userToken, validationCode)
	}

	handler := func(val interface{}) {
		val.(sender).send()
	}
	q := queue.NewQueue(handler, 20)

	q.Push(EmailInfo{To: userRegisterReq.Email, Code: validationCode})
	// spedisci via email il codice di validazione

	q.Wait()

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	err2 := json.NewEncoder(res).Encode(userRegisterRes)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

// UserCodeValidation validazione codice utente
func UserCodeValidation(res http.ResponseWriter, req *http.Request) {
	var userCodeValidationReq model.UserCodeValidationReq
	if req.Body == nil {
		http.Error(res, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&userCodeValidationReq)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	result := db.ValidateUserAppToken(userCodeValidationReq.ValidationCode, userCodeValidationReq.AppToken)
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	err2 := json.NewEncoder(res).Encode(model.ResHead{Success: result})
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
	// aggiorna USER_APP_TOKEN con informazione che utente ha validato email
	// CHECK_VALID => true
}

// Logout logout utente
func Logout(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	sToken := vars["token"]
	delete(Tokens, sToken)
	res.WriteHeader(http.StatusOK)
}

// Sessions elenco sessioni
func Sessions(res http.ResponseWriter, req *http.Request) {
	s := reflect.ValueOf(Tokens).MapKeys()
	var str string
	for _, element := range s {
		str += element.String() + ","
	}
	err := json.NewEncoder(res).Encode(model.Response{Data: str})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}
