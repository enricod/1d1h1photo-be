package rest

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/enricod/1h1dphoto.com-be/db"
	"github.com/enricod/1h1dphoto.com-be/model"
	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
	"github.com/otium/queue"
)

// ImgDownload client riceve jpg dell'immagine richiesta
func ImgDownload(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	imageid, _ := strconv.Atoi(vars["id"])
	submission, _ := db.SubmissionByID(uint(imageid))

	size := req.FormValue("size")

	var filename string
	if "t" == size {
		filename = model.Confs.ImgDir + "/" + submission.ImageUid + "_t"
	} else {
		filename = model.Confs.ImgDir + "/" + submission.ImageUid + "_s"
	}

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		// mandiamo immagine standard?
		panic(err)
	}

	resp.Header().Set("Content-Type", "image/jpg")
	resp.Write(f)
	//fmt.Printf("Req: %s %s %s\n", req.URL.Scheme, req.Host, req.URL.Path)
}

type imgProcess interface {
	process()
}

type imgInfo struct {
	ImagePath string
	OutDir    string
	ImageDim  uint
	ThumbDim  uint
}

func renamer(postfix string) model.RenamerType {
	return func(filename string) string {
		if strings.HasPrefix(filename, ".jpg") {
			return strings.Replace(filename, ".jpg", postfix+".jpg", 1)
		}
		return filename + postfix
	}
}

func toSquare() model.ImgTransform {
	return func(img image.Image, dim uint) (image.Image, error) {
		croppedImg, err := cutter.Crop(img, cutter.Config{
			Width:   4,
			Height:  4,
			Mode:    cutter.Centered,
			Options: cutter.Ratio,
		})
		m := resize.Resize(dim, 0, croppedImg, resize.Lanczos3)
		return m, err
	}
}

func toNewSize() model.ImgTransform {
	return func(img image.Image, dim uint) (image.Image, error) {
		m := resize.Resize(dim, 0, img, resize.Lanczos3)
		return m, nil
	}
}

func (ei imgInfo) process() {
	model.ImageManipulate(ei.ImagePath, ei.OutDir, ei.ImageDim, renamer("_s"), toNewSize())
	model.ImageManipulate(ei.ImagePath, ei.OutDir, ei.ImageDim, renamer("_t"), toSquare())
	log.Println("generazione thumbnails terminata  ", ei.ImagePath, ei.OutDir)
}

// ImgUpload caricamento di un'immagine dal telefono
func ImgUpload(res http.ResponseWriter, req *http.Request) {

	appToken := req.Header.Get("Authorization")

	// cercare un evento in corso, se esiste.
	eventoAperto := db.FindEventoByDate(time.Now())
	if eventoAperto == nil {
		// non esiste evento aperto, diamo errore al chiamante
		log.Println("nessun evento aperto, non posso accettare immagini")
		res.WriteHeader(http.StatusNotAcceptable)
		return
	}

	userAppToken := db.FindAppToken(appToken)
	file, header, err := req.FormFile("image")

	if err != nil {
		log.Println("[-] Error in req.FormFile ", err)
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, "{'error': %s}", err)
		return
	}
	defer file.Close()
	imageUID, _ := model.GenerateRandomString(32)

	errCreazioneDir := os.MkdirAll(model.Confs.ImgUploadDir, 0777)
	if errCreazioneDir != nil {
		log.Println("[-] Unable to create the directory for writing. Check your write access privilege.", errCreazioneDir)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	tempfile := model.Confs.ImgUploadDir + "/" + imageUID
	out, err := os.Create(tempfile)
	log.Println("tempfile ", tempfile)

	db.InsertSubmission(eventoAperto.ID, userAppToken, imageUID, header.Filename)
	if err != nil {
		log.Println("[-] Unable to create the file for writing. Check your write access privilege.", err)
		// fmt.Fprintf(res, "[-] Unable to create the file for writing. Check your write access privilege.", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
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

	//      RESIZE IMMAGINE
	handler := func(val interface{}) {
		val.(imgProcess).process()
	}
	q := queue.NewQueue(handler, 20)
	q.Push(imgInfo{ImagePath: tempfile, OutDir: model.Confs.ImgDir, ImageDim: 1024, ThumbDim: 250})
	q.Wait()

	// SCRIVI RESPONSE
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	respHead := model.ResHead{Success: true}
	err2 := json.NewEncoder(res).Encode(respHead)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}
