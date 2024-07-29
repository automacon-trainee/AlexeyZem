package main

import (
	"bytes"
	"reflect"
	"testing"
)

type TestCase struct {
	str   string
	wantR []rune
	wantB []byte
}

func TestFunc(t *testing.T) {
	tests := []TestCase{
		{"Hello, world!", []rune("Hello, world!"), []byte("Hello, world!")},
	}
	for _, test := range tests {
		resR := getRunes(test.str)
		if !reflect.DeepEqual(resR, test.wantR) {
			t.Errorf("getRunes(%q) = %v, want %v", test.str, resR, test.wantR)
		}
		resB := getBytes(test.str)
		if !bytes.Equal(resB, test.wantB) {
			t.Errorf("getBytes(%q) = %v, want %v", test.str, resB, test.wantB)
		}
	}
}
