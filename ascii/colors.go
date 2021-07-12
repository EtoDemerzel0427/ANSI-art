package ascii

import (
	"fmt"
	"image"
	"strings"
)

const reset = "\033[0m"
const sequence = "MESSI"
var seqNum = -1

func setForeColor(r, g, b uint32) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}


func setBackColor(r, g, b uint32) string {
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func nextChar() uint8{
	seqNum++
	if seqNum >= len(sequence) {
		seqNum = 0
	}
	return sequence[seqNum]
}

func Pixels2ColoredAscii(img image.Image) string {
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
				_, err := fmt.Fprintf(&sb, "%s%c", setForeColor(r, g, b), nextChar())
				if err != nil {
					return ""
				}
				oldr, oldg, oldb = r, g, b

			} else {
				if r == oldr && g == oldg && b == oldb {
					_, err := fmt.Fprintf(&sb, "%c", nextChar())
					if err != nil {
						return ""
					}
				} else {
					_, err := fmt.Fprintf(&sb, "%s%c", setForeColor(r, g, b), nextChar())
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

