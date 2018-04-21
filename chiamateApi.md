#registrazione user

POST /users/register

# request
{
  "username":"enricod",
  "email":"pippo@pluto.com"
}


# response

{
    "head": {
        "res": true
    },
    "body": {
        "userKey": "x8uGB",
        "user": {
            "id": 1,
            "username": "enricod",
            "email": "pippo@pluto.com"
        }
    }
}

===========================================================

#userCodeValidation
POST /users/codeValidation

#req
{
    ValidationCode: "ABCAS",
	AppToken:"123131276457135471571"
}

#res
code 200 se tutto ok


===========================================================

# elenco eventi
GET /events/summary/list

#req
nell'header deve esserci 
Authorization: <appToken>

#res
{
    futureEvents: [ 
        {
            Start       time.Time
	        End         time.Time
	        Name        string
	        Submissions:  [
                {
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
            ]
        }
    ],
    closedEvents: []

}
===========================================================


#logout user
/logout (fatto)

#session per i test
/sessions (fatto)

#registrazione user
/register

#informazioni user
/user/info

#cambia informazioni user
/user/info/update

#carica un'immagine, possibile tag per definire il giorno del caricamento foto
/image/insert

#carica le prime immagini per voti (valide)
/images/rank

#carica immagini per caricamento (valide)
/images/new

#carica le immagini del profilo (valide)
/images/profile/all

#immagini in caricamento (in validazione)
/images/profile/uploading

#vota immagine
/image/{imageId}/vote

#ritorna l'immagine
/image/{imageId}

Tutte le chiamate images tornano una lista di foto da vedere se tornarle solo gli id oppure tornarle con una risoluzione inferiore 


