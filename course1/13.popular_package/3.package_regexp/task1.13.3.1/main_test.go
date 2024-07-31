package main

import (
	"testing"
)

type TestCase struct {
	email string
	check bool
}

func TestIsValidEmail(t *testing.T) {
	testCases := []TestCase{
		{"abc@gmail.com", true},
		{"test@example.com", true},
		{"invalid_email", false},
	}
	for _, testCase := range testCases {
		if res := isValidEmail(testCase.email); res != testCase.check {
			t.Errorf("isValidEmail(%q) = %t; want %t", testCase.email, res, testCase.check)
		}
	}
}
