package main

import (
	"sort"
	"testing"
)

type TestCase struct {
	data []int
}

var tests = []TestCase{
	{[]int{5, 4, 3, 2, 1}},
	{[]int{5, 4, 3, 1, 3, 100, 4, 2}},
	{[]int{}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
	{[]int{15, 14, 13, 12, 11, 10, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
	{[]int{84, 42, 73, 63, 41, 37, 42, 7, 37, 38, 41, 92, 37, 88, 32, 10, 11, 60, 62,
		27, 48, 84, 33, 52, 8, 6, 4, 24, 4, 12, 80, 75, 15, 30, 40, 24, 31, 64, 24, 97, 20}},
}

func TestInsertionSort(t *testing.T) {
	for _, v := range tests {
		InsertionSort(v.data)
		if !sort.IntsAreSorted(v.data) {
			t.Errorf("InsertionSort(%v) failed", v.data)
		}
	}
}

func TestSelectionSortSort(t *testing.T) {
	for _, v := range tests {
		SelectionSort(v.data)
		if !sort.IntsAreSorted(v.data) {
			t.Errorf("SelectionSort(%v) failed", v.data)
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, v := range tests {
		res := MergeSort(v.data)
		if !sort.IntsAreSorted(res) {
			t.Errorf("MergeSort(%v) failed", res)
		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, v := range tests {
		QuickSort(v.data, 0, len(v.data)-1)
		if !sort.IntsAreSorted(v.data) {
			t.Errorf("QuickSort(%v) failed", v.data)
		}
	}
}

func TestGeneralSort(t *testing.T) {
	for _, v := range tests {
		GeneralSort(v.data)
		if !sort.IntsAreSorted(v.data) {
			t.Errorf("GeneralSort(%v) failed", v.data)
		}
	}
}
