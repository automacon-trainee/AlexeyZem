package main

import (
	"bytes"
	"testing"
)

type TestCase struct {
	str string
}

func TestGetReader(t *testing.T) {
	testCases := []TestCase{{str: "test"},
		{str: "test1"},
	}
	for _, testCase := range testCases {
		reader := getReader(bytes.NewBufferString(testCase.str))
		b := make([]byte, len(testCase.str))
		_, err := reader.Read(b)
		if err != nil {
			t.Fatal(err)
		}
		if string(b) != testCase.str {
			t.Errorf("expected: %s, got: %s", testCase.str, string(b))
		}
	}
}
