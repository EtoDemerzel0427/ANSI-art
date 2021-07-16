package decode

import (
	"ANSI-art/ansi"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
	"image/gif"
	"os"
	"time"
)

func Gif2imgs(filename string, GifWidth int, GifHeight int, duration time.Duration, seq string, blockMode bool) {
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

	for _, srcimg := range inGif.Image {
		img := image.NewNRGBA(rect)

		draw.Draw(img, srcimg.Bounds(), srcimg, srcimg.Rect.Min, draw.Src)
		img = imaging.Resize(img, GifWidth, GifHeight, imaging.Lanczos)

		fmt.Print(ansi.ClearScreen())
		if blockMode {
			fmt.Println(ansi.Pixels2ColoredBlocks(img))
		} else {
			fmt.Println(ansi.Pixels2ColoredANSI(img, seq))
		}

		time.Sleep(duration)

	}

}