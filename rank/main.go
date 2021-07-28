package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/EtoDemerzel0427/ANSI-art/ascii"
	"os"
	"reflect"
	"sort"
)

// Slice type wrapper for argsort
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

// writeLines writes the lines to the given file.
func writeLines(values interface{}, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	rv := reflect.ValueOf(values)
	if rv.Kind() != reflect.Slice {
		return errors.New("not a slice")
	}
	w := bufio.NewWriter(file)
	for i := 0; i < rv.Len(); i++ {
		fmt.Fprintln(w, rv.Index(i).Interface())
	}

	return w.Flush()
}

func main() {
	darkness := make([]int, 95)

	for i := 32; i <= 126; i++ {
		darkness[i-32] = ascii.DrawFont(rune(i), "../font/monaco.ttf", 100, 72, "none")
	}


	s := NewSlice(darkness...)
	sort.Sort(s)
	//fmt.Println(s.IntSlice)
	intensity := make([]float64, 95)
	for i, v := range s.IntSlice {
		intensity[i] = float64(v - s.IntSlice[0]) / float64(s.IntSlice[94] - s.IntSlice[0]) * 128
	}

	fmt.Println(intensity)
	for i, v := range s.idx {
		s.idx[i] = v + 32
		fmt.Printf("%c ", v + 32)
	}
	writeLines(intensity, "intensity.txt")
	writeLines(s.idx, "rank.txt")
	fmt.Println()
}
