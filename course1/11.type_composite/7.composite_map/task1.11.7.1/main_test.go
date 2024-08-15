package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	text string
	want map[string]int
}

func TestCountWordOccurrences(t *testing.T) {
	tests := []TestCase{
		{"a b c b c c", map[string]int{"a": 1, "b": 2, "c": 3}},
		{"Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum",
			map[string]int{"Lorem": 1, "ipsum": 2, "dolor": 1, "sit": 1, "consectetur": 1, "adipiscing": 1, "elit": 1, "amet": 1}},
	}
	for _, test := range tests {
		res := CountWordOccurrences(test.text)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("TestCountWordOccurrences(%s)=%v, want %v", test.text, res, test.want)
		}
	}
}
