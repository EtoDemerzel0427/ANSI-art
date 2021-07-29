package main

import (
	"fmt"
	"sort"
)

func main() {
	darkness := make([]int, 95)

	for i := 32; i <= 126; i++ {
		darkness[i-32] = drawFont(rune(i), "../font/monaco.ttf", 100, 72, "none")
	}

	s := NewSlice(darkness...)
	sort.Sort(s)
	//fmt.Println(s.IntSlice)
	intensity := make([]float64, 95)
	for i, v := range s.IntSlice {
		intensity[i] = float64(v-s.IntSlice[0]) / float64(s.IntSlice[94]-s.IntSlice[0]) * 128
	}

	fmt.Println(intensity)
	for i, v := range s.idx {
		s.idx[i] = v + 32
		fmt.Printf("%c ", v+32)
	}
	writeLines(intensity, "intensity.txt")
	writeLines(s.idx, "rank.txt")
	fmt.Println()
}
