package main

import (
	"testing"
)

type TestCase struct {
	text string
	old  rune
	newR rune
	want string
}

func TestReplaceSymbols(t *testing.T) {
	testCases := []TestCase{
		{"Hello, world!", 'o', '0', "Hell0, w0rld!"},
	}
	for _, testCase := range testCases {
		res := ReplaceSymbols(testCase.text, testCase.old, testCase.newR)
		if res != testCase.want {
			t.Errorf("replaceSymbols failed. text: %s, old: %c, new: %c, want: %s",
				testCase.text, testCase.old, testCase.newR, testCase.want)
		}
	}
}
