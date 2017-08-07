package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/enricod/1h1dphoto.com-be/routes"
)

func main() {
	port := flag.String("port", "9090", "porta servizio")
	flag.Parse()

	log.Printf("Avvio server su porta %v\n", *port)
	log.Fatal(http.ListenAndServe(":" + *port, routes.NewRouter()))
}
