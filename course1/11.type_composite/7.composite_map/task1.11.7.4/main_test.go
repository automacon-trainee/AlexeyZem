package main

import (
	"testing"
)

type TestCase struct {
	text   string
	filter map[string]bool
	want   string
}

func TestFilterSentence(t *testing.T) {
	testCases := []TestCase{
		{"Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum",
			map[string]bool{"ipsum": true, "consectetur": true, "elit": true},
			"Lorem dolor sit amet adipiscing"},
	}
	for _, testCase := range testCases {
		res := filterSentence(testCase.text, testCase.filter)
		if res != testCase.want {
			t.Errorf("filterSentence(%q, %v) = %q, want %q", testCase.text, testCase.filter, res, testCase.want)
		}
	}
}
