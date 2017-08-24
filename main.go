package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/enricod/1h1dphoto.com-be/routes"
	"github.com/enricod/1h1dphoto.com-be/rest"
)

func main() {
	port 		:= flag.String("port", "9090", "porta servizio")
	imagesDir 	:= flag.String("imgDir", "/tmp/images", "directory per immagini finali")
	imagesUploadDir 	:= flag.String("imgUploadDir", "/tmp/imagesTemp", "directory per immagini caricate da utente")
	rest.ImagesDir = *imagesDir
	rest.ImagesUploadDir = *imagesUploadDir
	flag.Parse()

	log.Printf("Avvio server su porta %v\n", *port)
	log.Printf("Images dir %v\n", *imagesDir)
	log.Printf("Images upload dir %v\n", *imagesUploadDir)
	log.Fatal(http.ListenAndServe(":" + *port, routes.NewRouter()))
}
