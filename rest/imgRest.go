package rest

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"io/ioutil"
	"log"
	"fmt"
	"github.com/enricod/1h1dphoto.com-be/model"
	"os"
	"io"
	"github.com/enricod/1h1dphoto.com-be/db"
	"encoding/json"
)

func ImgDownload(resp http.ResponseWriter, req *http.Request) {
	//fmt.Println(req.URL.RequestURI())
	vars := mux.Vars(req)
	imageid,_ := strconv.Atoi( vars["id"])
	submission, _ := db.SubmissionById( uint(imageid) )

	size := req.FormValue("size")

	var filename string
	if "t" == size {
		filename = ImagesDir + "/" + submission.ImageUid + "_t"
	} else {
		filename = ImagesDir + "/" + submission.ImageUid + "_s"
	}
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	resp.Header().Set("Content-Type", "image/jpg")
	resp.Write(f)
	//fmt.Printf("Req: %s %s %s\n", req.URL.Scheme, req.Host, req.URL.Path)
}


func ImgUpload(res http.ResponseWriter, req *http.Request) {

	appToken := req.Header.Get("Authorization")

	// FIXME
	eventId := uint(1)

	userAppToken := db.FindAppToken(appToken)

	log.Println(appToken)
	file, header, err := req.FormFile("image")

	if err != nil {
		log.Println("[-] Error in req.FormFile ", err)
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, "{'error': %s}", err)
		return
	}
	defer file.Close()
	imageUid, _ := model.GenerateRandomString(32)

	out, err := os.Create(   "/tmp/" + imageUid)
	os.Rename("/tmp/" + imageUid, ImagesUploadDir +"/"+ imageUid)

	db.InsertSubmission(eventId, userAppToken, imageUid, header.Filename)
	if err != nil {
		log.Println("[-] Unable to create the file for writing. Check your write access privilege.", err)
		fmt.Fprintf(res, "[-] Unable to create the file for writing. Check your write access privilege.", err)
		res.WriteHeader(http.StatusInternalServerError)
	}
	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println("[-] Error copying file.", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("[+] File uploaded successfully: uploaded-", header.Filename)

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)

	respHead := model.ResHead{Success:true}
	err2 := json.NewEncoder(res).Encode(respHead)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}
