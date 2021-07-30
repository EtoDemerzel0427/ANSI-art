package art

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"log"
	"strings"
)

type Mode int

const (
	asciiText Mode = iota
	ansiText
	ansiBlock
)

type Solver struct {
	intensity []float64
	rank      []int
	width     int
	height    int
	contrast  float64
	sigma     float64
	seq       string
	mode      Mode
	Convert   func(image2 image.Image) string
}

func NewSolver(width, height int, contrast, sigma float64, seq string, mode Mode) (as *Solver) {
	// the intensity/rank files are predefined
	intensity, _ := readFloatLines(Root + "/rank/intensity.txt")
	rank, _ := readIntLines(Root + "/rank/rank.txt")

	as = &Solver{
		intensity: intensity,
		rank:      rank,
		width:     width,
		height:    height,
		contrast:  contrast,
		sigma:     sigma,
		seq:       seq,
		mode:      mode,
		Convert:   nil,
	}

	switch as.mode {
	case asciiText:
		as.Convert = as.pixels2Ascii
	case ansiText:
		as.Convert = as.pixels2ColoredANSI
	case ansiBlock:
		as.Convert = as.pixels2ColoredBlocks
	default:
		log.Fatal("undefined mode")
	}

	return
}

func (as *Solver) TuneImage(src image.Image) *image.NRGBA {
	dst := imaging.Resize(src, as.width, as.height, imaging.Lanczos)
	dst = imaging.AdjustContrast(dst, as.contrast)
	dst = imaging.Sharpen(dst, as.sigma)
	if as.mode == asciiText {
		dst = imaging.Grayscale(dst)
	}

	return dst
}

func (as *Solver) pixels2Ascii(img image.Image) string {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	sb := strings.Builder{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, _, _, _ := img.At(x, y).RGBA() // r, g, b are the same for grayscale image
			r >>= 8
			asciiIdx := findClosestK(int(r), as.intensity)
			c := as.rank[asciiIdx]

			_, err := fmt.Fprintf(&sb, "%c", c)
			if err != nil {
				return ""
			}
		}
		_, fprintln := fmt.Fprintln(&sb)
		if fprintln != nil {
			return ""
		}
	}

	return sb.String()
}

func (as *Solver) pixels2ColoredANSI(img image.Image) string {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	sb := strings.Builder{}
	_, err := fmt.Fprintf(&sb, setBackColor(0, 0, 0))
	if err != nil {
		return ""
	}

	var oldr, oldg, oldb uint32
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r >>= 8
			g >>= 8
			b >>= 8

			if x == 0 && y == 0 {
				_, err := fmt.Fprintf(&sb, "%s%c", setForeColor(r, g, b), nextChar(as.seq))
				if err != nil {
					return ""
				}
				oldr, oldg, oldb = r, g, b

			} else {
				if r == oldr && g == oldg && b == oldb {
					_, err := fmt.Fprintf(&sb, "%c", nextChar(as.seq))
					if err != nil {
						return ""
					}
				} else {
					_, err := fmt.Fprintf(&sb, "%s%c", setForeColor(r, g, b), nextChar(as.seq))
					if err != nil {
						return ""
					}
				}
				oldr, oldg, oldb = r, g, b
			}
		}
		_, err := fmt.Fprintln(&sb)
		if err != nil {
			return ""
		}
	}
	_, err = fmt.Fprintln(&sb, reset)
	if err != nil {
		return ""
	}

	return sb.String()
}

func (as *Solver) pixels2ColoredBlocks(img image.Image) string {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	sb := strings.Builder{}

	var oldr, oldg, oldb uint32
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r >>= 8
			g >>= 8
			b >>= 8

			if x == 0 && y == 0 {
				_, err := fmt.Fprintf(&sb, "%s ", setBackColor(r, g, b))
				if err != nil {
					return ""
				}
				oldr, oldg, oldb = r, g, b

			} else {
				if r == oldr && g == oldg && b == oldb {
					_, err := fmt.Fprint(&sb, " ")
					if err != nil {
						return ""
					}
				} else {
					_, err := fmt.Fprintf(&sb, "%s ", setBackColor(r, g, b))
					if err != nil {
						return ""
					}
				}
				oldr, oldg, oldb = r, g, b
			}
		}
		_, err := fmt.Fprintln(&sb)
		if err != nil {
			return ""
		}
	}
	_, err := fmt.Fprintln(&sb, reset)
	if err != nil {
		return ""
	}

	return sb.String()
}
