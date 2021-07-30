package art

import (
	"fmt"
)

const reset = "\033[0m"

//const sequence = "01"
var seqNum = -1

func MoveCursor(row, col int) string {
	return fmt.Sprintf("\033[%d;%dH", row+1, col+1) // in ansi the coordinators start from (1,1)
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

func nextChar(sequence string) uint8 {
	seqNum++
	if seqNum >= len(sequence) {
		seqNum = 0
	}
	return sequence[seqNum]
}
