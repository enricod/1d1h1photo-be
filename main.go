package main

import (
	"log"
	"net/http"

	"1d1hphoto.com-be/routes"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", routes.NewRouter()))
}
