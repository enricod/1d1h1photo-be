package model

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"
)

type RenamerType func(string) string

type ImgTransform func(image2 image.Image, dim uint) (image.Image, error)

func NomeImmagine(absolutePath string) string {
	splitted := strings.Split(absolutePath, "/")
	return splitted[len(splitted)-1]
}

func ReadImageFromFileSystem(fileName string) (image.Image, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	img, err := jpeg.Decode(file)
	file.Close()
	return img, err
}

func ImageManipulate(fileName string, outDir string, dim uint, renamer RenamerType, operation ImgTransform) {
	nomeImmagine := NomeImmagine(fileName)
	img, err := ReadImageFromFileSystem(fileName)
	if err != nil {
		log.Fatal(err)
	}
	m, err := operation(img, dim)
	out, err := os.Create(outDir + "/" + renamer(nomeImmagine))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
