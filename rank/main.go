package main

import (
	"ANSI-art/ascii"
	"fmt"
	"sort"
)

type Slice struct {
	sort.IntSlice
	idx []int
}

func (s Slice) Swap(i, j int) {
	s.IntSlice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func NewSlice(n ...int) *Slice {
	s := &Slice{IntSlice: sort.IntSlice(n), idx: make([]int, len(n))}
	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}

func main() {
	darkness := make([]int, 95)

	for i := 32; i <= 126; i++ {
		darkness[i-32] = ascii.DrawFont(rune(i), "../font/monaco.ttf", 100, 72, "none")
	}


	s := NewSlice(darkness...)
	sort.Sort(s)
	//fmt.Println(s.IntSlice)
	for _, v := range s.idx {
		fmt.Printf("%c ", v + 32)
	}
	fmt.Println()
}
