package main

import (
	"bytes"
	"strings"
	"testing"
)

type TestCase struct {
	data []byte
}

func TestGetScanner(t *testing.T) {
	testCases := []TestCase{
		{[]byte("hello world")},
		{[]byte("hello\nworld")},
	}
	for _, testCase := range testCases {
		scanner := getScanner(bytes.NewBuffer(testCase.data))
		if scanner == nil {
			t.Errorf("scanner is nil")
		}
		text := ""
		for scanner.Scan() {
			text += scanner.Text()
			text += "\n"
		}
		text = strings.TrimSuffix(text, "\n")
		if text != string(testCase.data) {
			t.Errorf("scanner text is wrong. expected: %s, got: %s", testCase.data, text)
		}
	}
}
