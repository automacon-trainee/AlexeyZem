package main

import (
	"testing"
)

type testCaseSum struct {
	data   [8]int
	answer int
}

func TestSum(t *testing.T) {
	tests := []testCaseSum{
		{data: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, answer: 36},
		{data: [8]int{0, 0, 0, 0, 0, 0, 0, 0}, answer: 0},
	}
	for _, v := range tests {
		answer := sum(v.data)
		if answer != v.answer {
			t.Errorf("sum(%d) = %d, want %d", v.data, answer, v.answer)
		}
	}
}

type testCaseAverage struct {
	data   [8]int
	answer float64
}

func TestAverage(t *testing.T) {
	tests := []testCaseAverage{
		{data: [8]int{5, 2, 3, 4, 5, 6, 7, 8}, answer: 5.0},
		{data: [8]int{0, 0, 0, 0, 0, 0, 0, 0}, answer: 0.0},
	}
	for _, v := range tests {
		answer := average(v.data)
		if answer != v.answer {
			t.Errorf("average(%d) = %f, want %f", v.data, answer, v.answer)
		}
	}
}

type TestCaseAverageFloat struct {
	data   [8]float64
	answer float64
}

func TestAverageFloat(t *testing.T) {
	tests := []TestCaseAverageFloat{
		{data: [8]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}, answer: 4.5},
		{data: [8]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}, answer: 0.0},
	}
	for _, v := range tests {
		answer := averageFloat(v.data)
		if answer != v.answer {
			t.Errorf("average(%f) = %f, want %f", v.data, answer, v.answer)
		}
	}
}

type TestCaseReverse struct {
	data   [8]int
	answer [8]int
}

func TestReverse(t *testing.T) {
	tests := []TestCaseReverse{
		{data: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, answer: [8]int{8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, v := range tests {
		answer := reverse(v.data)
		if answer != v.answer {
			t.Errorf("reverse(%d) = %d, want %d", v.data, answer, v.answer)
		}
	}
}
