package decode

import (
	"fmt"
	"github.com/EtoDemerzel0427/ANSI-art/art"
	"image"
	"image/draw"
	"image/gif"
	"os"
	"time"
)

type GifDecoder struct {
	as *art.Solver
	loopNum int
	duration time.Duration
	music *chan bool
}

func NewGifDecoder(as *art.Solver, loopNum int, duration time.Duration, music *chan bool) *GifDecoder {
	return &GifDecoder{
		as: as,
		loopNum: loopNum,
		duration: duration,
		music: music,
	}

}

func (gd *GifDecoder) Decode(filename string) {
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

	for i := 0; i < gd.loopNum; i++ {
		for _, srcimg := range inGif.Image {
			img := image.NewNRGBA(rect)

			draw.Draw(img, srcimg.Bounds(), srcimg, srcimg.Rect.Min, draw.Src)
			img = gd.as.TuneImage(img)

			fmt.Print(art.ClearScreen())
			fmt.Println(gd.as.Convert(img))
			//if asciiMode {
			//	img = imaging.Grayscale(img)
			//	fmt.Println(art.Pixels2Ascii(img))
			//
			//} else {
			//	if blockMode {
			//		fmt.Println(art.Pixels2ColoredBlocks(img))
			//	} else {
			//		fmt.Println(art.Pixels2ColoredANSI(img, seq))
			//	}
			//}
			time.Sleep(gd.duration)
		}
	}
	*(gd.music) <- true
}