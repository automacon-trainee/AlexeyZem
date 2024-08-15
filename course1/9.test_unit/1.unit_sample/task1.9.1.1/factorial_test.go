package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	{
		result := Factorial(0)
		if result != 1 {
			t.Errorf("TestFactorial failed.")
		}
	}

	{
		result := Factorial(1)
		if result != 1 {
			t.Errorf("TestFactorial failed.")
		}
	}

	{
		result := Factorial(5)
		if result != 120 {
			t.Errorf("TestFactorial failed.")
		}
	}
}
