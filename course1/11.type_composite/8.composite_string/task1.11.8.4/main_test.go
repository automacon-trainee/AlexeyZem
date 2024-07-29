package main

import (
	"testing"
)

type TestCase struct {
	str  []string
	want string
}

func TestConcatStrings(t *testing.T) {
	tests := []TestCase{
		{[]string{"a", " ", "b"}, "a b"},
	}
	for _, test := range tests {
		res := concatStrings(test.str...)
		if res != test.want {
			t.Errorf("concatStrings(%v) got %s, want %s", test.str, res, test.want)
		}
	}
}
