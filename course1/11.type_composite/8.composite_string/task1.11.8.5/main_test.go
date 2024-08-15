package main

import (
	"testing"
)

type TestCase struct {
	str  string
	want string
}

func TestReverseString(t *testing.T) {
	tests := []TestCase{
		{str: "hello world", want: "dlrow olleh"},
	}
	for _, test := range tests {
		res := ReverseString(test.str)
		if res != test.want {
			t.Errorf("ReverseString(%q) failed. Got %q, want %q", test.str, res, test.want)
		}
	}
}
