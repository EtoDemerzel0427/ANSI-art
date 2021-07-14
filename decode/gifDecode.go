package decode

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
	"path"
	"strings"
)

func Gif2imgs(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		_, err1 := fmt.Fprintln(os.Stderr, err)
		if err1 != nil {
			return
		}
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	inGif, err := gif.DecodeAll(f)
	if err != nil {
		_, err1 := fmt.Fprintln(os.Stderr, err)
		if err1 != nil {
			return
		}
	}

	config, _ := gif.DecodeConfig(f)
	rect := image.Rect(0, 0, config.Width, config.Height)
	if rect.Min == rect.Max {
		var max image.Point
		for _, frame := range inGif.Image {
			maxF := frame.Bounds().Max
			if max.X < maxF.X {
				max.X = maxF.X
			}
			if max.Y < maxF.Y {
				max.Y = maxF.Y
			}
		}
		rect.Max = max
	}

	format := fmt.Sprintf("%%0%dd", len(string(rune(len(inGif.Image))))+1)
	for i, srcimg := range inGif.Image {
		img := image.NewRGBA(rect)
		if _, err := os.Stat(strings.Split(filename, ".")[0]); os.IsNotExist(err) {
			err := os.Mkdir(strings.Split(filename, ".")[0], os.ModePerm)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}

		subfn := path.Join(strings.Split(filename, ".")[0], fmt.Sprintf(format, i) + ".png")
		//subfn := strings.Split(filename, ".")[0] + fmt.Sprintf(format, i) + ".png"
		f1, err := os.Create(subfn)
		if err != nil {
			panic(err)
		}

		draw.Draw(img, srcimg.Bounds(), srcimg, srcimg.Rect.Min, draw.Src)
		fmt.Printf("\r%s", subfn)
		err = png.Encode(f1, img)
		if err != nil {
			_, err1 := fmt.Fprintln(os.Stderr, err)
			if err1 != nil {
				return
			}
			return
		}
		err = f1.Close()
		if err != nil {
			_, err1 := fmt.Fprintln(os.Stderr, err)
			if err1 != nil {
				return
			}
			return
		}

	}
}