package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	str  string
	want map[rune]int
}

func TestCountRussianLetters(t *testing.T) {
	tests := []TestCase{
		{"Привет мир!", map[rune]int{
			'П': 1,
			'р': 2,
			'и': 2,
			'в': 1,
			'е': 1,
			'т': 1,
			'м': 1,
		}},
	}
	for _, tt := range tests {
		got := countRussianLetters(tt.str)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("countRussianLetters(%v) got %v, want %v", tt.str, got, tt.want)
		}
	}
}
