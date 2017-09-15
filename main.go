package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/enricod/1h1dphoto.com-be/model"

	"github.com/enricod/1h1dphoto.com-be/routes"
)

func defaultConfs() model.AppConfs {
	return model.AppConfs{Port: 9090, ImgDir: "/tmp/images", ImgUploadDir: "/tmp/imagesTemp"}
}

func main() {
	conffile := flag.String("c", "/etc/1h1d.json", "file configurazione")
	flag.Parse()

	// oggetto con configurazioni
	var confs model.AppConfs

	log.Printf("lettura configurazioni da file %v\n", *conffile)
	content, err := ioutil.ReadFile(*conffile)
	if err != nil {
		confs = defaultConfs()
	} else {
		err2 := json.Unmarshal(content, &confs)
		if err2 != nil {
			//Do something
			log.Printf("Errore in lettura file configurazione", *conffile)
			confs = defaultConfs()
		}
	}

	model.Confs = confs

	log.Printf("Avvio server su porta %v\n", confs.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(confs.Port), routes.NewRouter()))
}
