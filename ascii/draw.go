package ascii

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"os"
)

func DrawFont(text rune, fontFile string, fontSize float64, dpi float64, hinting string) int {
	data, err := os.ReadFile(fontFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	f, _ := truetype.Parse(data)

	fg, bg := image.Black, image.White
	rgba := image.NewRGBA(image.Rect(0, 0, 250, 200))
	draw.Draw(rgba, rgba.Bounds(), bg, image.Point{}, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	switch hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	ruler := color.RGBA{R: 0xdd, G: 0xdd, B: 0xdd, A: 0xff}

	for i := 0; i < 200; i++ {
		rgba.Set(0, i, ruler)
	}

	// Truetype stuff
	opts := truetype.Options{}
	opts.Size = 125.0
	face := truetype.NewFace(f, &opts)

	// Calculate the widths and print to image

	awidth, _ := face.GlyphAdvance(text)

	iwidthf := int(float64(awidth) / 64)
	pt := freetype.Pt(125-iwidthf/2, 128)
	_, _ = c.DrawString(string(text), pt)

	size := rgba.Bounds().Size()
	blackCnt := 0
	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			pixel := rgba.At(i, j)
			r, g, b, a := pixel.RGBA()
			r >>= 8
			g >>= 8
			b >>= 8
			a >>= 8
			if r != 255 || g != 255 || b != 255 {
				blackCnt++
			}
		}
	}

	return blackCnt
}
