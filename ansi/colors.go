package ansi

import (
	"fmt"
	"image"
	"strings"
)

const reset = "\033[0m"
//const sequence = "01"
var seqNum = -1

func MoveCursor(row, col int) string {
	return fmt.Sprintf("\033[%d;%dH", row+1, col+1)  // in ansi the coordinators start from (1,1)
}

func ClearScreen() string {
	s := fmt.Sprint("\033[2J")
	s += MoveCursor(0, 0)
	return s
}

func setForeColor(r, g, b uint32) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}


func setBackColor(r, g, b uint32) string {
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func nextChar(sequence string ) uint8{
	seqNum++
	if seqNum >= len(sequence) {
		seqNum = 0
	}
	return sequence[seqNum]
}

func Pixels2ColoredANSI(img image.Image, seq string) string {
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
				_, err := fmt.Fprintf(&sb, "%s%c", setForeColor(r, g, b), nextChar(seq))
				if err != nil {
					return ""
				}
				oldr, oldg, oldb = r, g, b

			} else {
				if r == oldr && g == oldg && b == oldb {
					_, err := fmt.Fprintf(&sb, "%c", nextChar(seq))
					if err != nil {
						return ""
					}
				} else {
					_, err := fmt.Fprintf(&sb, "%s%c", setForeColor(r, g, b), nextChar(seq))
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

func Pixels2ColoredBlocks(img image.Image) string {
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
					_, err := fmt.Fprint(&sb,  " ")
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

