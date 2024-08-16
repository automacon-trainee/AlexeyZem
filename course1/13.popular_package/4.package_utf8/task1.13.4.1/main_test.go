package main

import (
	"testing"
)

type TestCase struct {
	str  string
	want int
}

func TestCountUniqueUTF8Characters(t *testing.T) {
	testCases := []TestCase{
		{"Hello,	!", 7},
		{"aaaaaaaaaaaaaaaa", 1},
	}
	for _, testCase := range testCases {
		got := countUniqueUTF8Characters(testCase.str)
		if got != testCase.want {
			t.Errorf("Got %d, want %d", got, testCase.want)
		}
	}

}
