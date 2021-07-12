package main

import (
	"fmt"
	"image"
	"log"
	"strings"

	"github.com/disintegration/imaging"
)

var pixels = []uint8(" .,:;i1tfLCG08@")

func pixels2Ascii(img image.Image) string {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	b := strings.Builder{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grayScale, _, _, _ := img.At(x, y).RGBA()
			grayScale >>= 8
			_, err := fmt.Fprintf(&b, "%c ", pixels[grayScale/18])
			if err != nil {
				return ""
			}


		}
		_, err := fmt.Fprintln(&b)
		if err != nil {
			return ""
		}
	}

	return b.String()
}


func main() {
	// Open a test image.
	src, err := imaging.Open("/Users/weiran/Desktop/CodeSpace/GolangCode/ANSI-art/pic/wiki.jpeg")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Crop the original image to 300x300px size using the center anchor.
	src = imaging.CropAnchor(src, 300, 300, imaging.Center)

	// Resize the cropped image to width = 200px preserving the aspect ratio.
	src = imaging.Resize(src, 50, 0, imaging.Lanczos)
	src = imaging.Grayscale(src)

	fmt.Print(pixels2Ascii(src))




}





