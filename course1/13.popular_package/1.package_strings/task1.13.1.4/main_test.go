package main

import (
	"testing"
)

func TestGenerateActivationKey(t *testing.T) {
	result1 := generateActivationKey()
	result2 := generateActivationKey()

	if len(result1) != len(result2) || len(result1) != 19 || result1 == result2 {
		t.Errorf("generateActivationKey(): len should be 19, but results should be not equal")
	}
}
