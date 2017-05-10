package main

import (
	"./routes"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", routes.NewRouter()))
}
