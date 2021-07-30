package art

import (
	"bufio"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../")
)

// readIntLines reads a whole file into memory
// and returns a slice of its lines.
func readIntLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intVal, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, intVal)
	}
	return lines, scanner.Err()
}

func readFloatLines(path string) ([]float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intVal, _ := strconv.ParseFloat(scanner.Text(), 64)
		lines = append(lines, intVal)
	}
	return lines, scanner.Err()
}

func findClosestK(value int, arr []float64) (id int) {
	x := float64(value)
	low, high, mid := 0, len(arr)-1, 0
	for low < high {
		mid = low + (high-low)/2
		if arr[mid] == x {
			// this is very unlikely for floats, but we still keep this
			return mid
		}

		if x < arr[mid] {
			if mid > 0 && x > arr[mid-1] {
				d1 := math.Abs(arr[mid] - x)
				d2 := math.Abs(arr[mid-1] - x)
				if d1 < d2 {
					return mid
				} else {
					return mid - 1
				}
			}
			high = mid

		} else {
			if mid < len(arr)-1 && x < arr[mid+1] {
				d1 := math.Abs(arr[mid] - x)
				d2 := math.Abs(arr[mid+1] - x)
				if d1 < d2 {
					return mid
				} else {
					return mid + 1
				}
			}
			low = mid + 1
		}

	}
	return low
}
