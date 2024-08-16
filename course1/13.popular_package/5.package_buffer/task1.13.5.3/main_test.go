package main

import (
	"bytes"
	"testing"
)

type TestCase struct {
	str string
}

func TestGetDataString(t *testing.T) {
	testCases := []TestCase{
		{str: "hello world"},
		{str: "test 1"},
		{str: "test 2"},
		{str: "test 3"},
		{str: "test 4"},
		{str: "Привет мир"},
		{str: ""},
		{str: "123"},
		{str: "ТЕстовые данные для проверки РабоТы функции GetDATAsTriNg"},
	}
	for _, testCase := range testCases {
		got := getDataString(bytes.NewBufferString(testCase.str))
		if got != testCase.str {
			t.Errorf("getDataString(%v) = %v, want %v", testCase.str, got, testCase.str)
		}
	}
}
