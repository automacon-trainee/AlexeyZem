package main

import (
	"testing"
)

type TestCase struct {
	text     string
	wantSym  int
	wantByte int
}

func TestFunc(t *testing.T) {
	tests := []TestCase{
		{"Привет, мир!", 12, 21},
	}
	for _, test := range tests {
		sym := countSymbols(test.text)
		byt := countBytes(test.text)
		if sym != test.wantSym {
			t.Errorf("countSymbols failed: got %d, want %d", sym, test.wantSym)
		}
		if byt != test.wantByte {
			t.Errorf("countBytes failed: got %d, want %d", byt, test.wantByte)
		}
	}
}
