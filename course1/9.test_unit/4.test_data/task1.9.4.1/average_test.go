package main

import (
	"math/rand"
	"reflect"
	"testing"
)

type testCase struct {
	arr    []float64
	answer float64
}

func TestAverage(t *testing.T) {
	tests := []testCase{
		{[]float64{1, 2, 3, 4, 5}, 3},
		{[]float64{1}, 1},
		{[]float64{}, 0},
		{[]float64{0, 1, 2}, 1},
		{[]float64{2, 3}, 2.5},
	}
	for _, test := range tests {
		res := average(test.arr)
		if res != test.answer {
			t.Errorf("average(%+v) => %f, want %f", test.arr, res, test.answer)
		}
	}

	firstSl := generateSlice(10)
	secondSl := generateSlice(10)
	if reflect.DeepEqual(firstSl, secondSl) {
		t.Errorf("Not random generated slice")
	}
}

func generateSlice(size int) []float64 {
	res := make([]float64, size)
	for i := 0; i < size; i++ {
		res[i] = rand.Float64()
	}
	return res
}
