package ascii

import (
	"math"
	"testing"
)

func TestReadFloatLines(t *testing.T) {
	intensity, err := ReadFloatLines("/Users/weiran/Desktop/CodeSpace/GolangCode/ANSI-art/rank/intensity.txt")
	if err != nil {
		t.Error(err)
	}

	if len(intensity) != 95 {
		t.Error("length is incorrect")
	}

	if math.Abs(intensity[0]) > 1e-7 {
		t.Error("first value incorrect")
	}

	if math.Abs(intensity[len(intensity)-1] - 128) > 1e-7 {
		t.Error("last value incorrect")
	}

	if math.Abs(intensity[10] - 33.85806451612903) > 1e-7 {
		t.Error("tenth value incorrect")
	}
}

func TestReadIntLines(t *testing.T) {
	rank, err := ReadIntLines("/Users/weiran/Desktop/CodeSpace/GolangCode/ANSI-art/rank/rank.txt")
	if err != nil {
		t.Error(err)
	}

	if len(rank) != 95 {
		t.Error("length is incorrect")
	}

	if rank[0] != 32 {
		t.Error("first value incorrect")
	}

	if rank[len(rank)-1] != 87 {
		t.Error("last value incorrect")
	}

	if rank[10] != 124 {
		t.Error("tenth value incorrect")
	}
}

func TestFindClosestK(t *testing.T) {
	x := []float64{1., 2., 100., 100., 100., 120.}
	if FindClosestK(2, x) != 1 {
		t.Errorf("wrong pos:%d, should be 1", FindClosestK(2, x))
	}

	x = []float64{
		9.703225806451613,
		11.664516129032258,
		17.651612903225807,
		20.18064516129032,
		22.296774193548387,
		23.12258064516129,
		23.948387096774194,
		29.212903225806453,
		29.316129032258065,
		33.85806451612903,
		35.66451612903226,
		39.89677419354839,
		42.89032258064516,
	}
	if FindClosestK(23, x) != 5 {
		t.Errorf("wrong pos:%d, should be 5", FindClosestK(23, x))
	}

	if FindClosestK(37, x) != 10 {
		t.Errorf("wrong pos:%d, should be 10", FindClosestK(37, x))
	}
	if FindClosestK(39, x) != 11 {
		t.Errorf("wrong pos:%d, should be 10", FindClosestK(39, x))
	}
	if FindClosestK(0, x) != 0 {
		t.Errorf("wrong pos:%d, should be 0", FindClosestK(0, x))
	}

	if FindClosestK(100, x) != 12 {
		t.Errorf("wrong pos:%d, should be 12", FindClosestK(100, x))
	}
}
