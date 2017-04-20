package main

import (
	"it/1d1hphoto/routes"
	"log"

	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", routes.NewRouter()))
}
