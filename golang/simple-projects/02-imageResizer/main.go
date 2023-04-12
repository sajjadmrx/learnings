package main

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	logoPath := "./dns-logo.png"
	resizer(200, 200, &logoPath)
}

func resizer(w uint, h uint, path *string) {

	file, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	m := resize.Resize(w, h, img, resize.Lanczos3)
	out, err := os.Create("resized.png")
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	jpeg.Encode(out, m, nil)

}
