package main

import (
	"fmt"
	"testing"
	"time"
)

func CompareWithFactorialFaster() map[string]bool {
	res := make(map[string]bool)
	resTimeI := 0 * time.Second
	resTimeR := 0 * time.Second
	for i := 10; i <= 100; i += 10 {
		startR := time.Now()
		_ = FactorialRecursive(i)
		resTimeR += time.Since(startR)
		startI := time.Now()
		_ = FactorialIterative(i)
		resTimeI += time.Since(startI)
	}
	if resTimeI < resTimeR {
		res["Factorial Iterative"] = true
		res["Factorial Recursive"] = false
	} else {
		res["Factorial Iterative"] = false
		res["Factorial Recursive"] = true
	}
	return res
}

func TestInSmallData(t *testing.T) {
	fmt.Println(CompareWithFactorialFaster())
}

func TestInBigData(t *testing.T) {
	startR := time.Now()
	_ = FactorialRecursive(100000)
	elapsedR := time.Since(startR)
	startI := time.Now()
	_ = FactorialIterative(100000)
	elapsedI := time.Since(startI)
	if elapsedR > elapsedI {
		fmt.Printf("Factorial Iterative faster in bigData. Time recursive:%v, Time Iterative:%v\n", elapsedR, elapsedI)
	} else {
		fmt.Printf("Factorial Recursive faster in BigData. Time recursive:%v, Time Iterative:%v\n", elapsedR, elapsedI)
	}
}

type TestCase struct {
	data int
	want int
}

func TestFactorialIterative(t *testing.T) {
	tests := []TestCase{
		{data: 1, want: 1},
		{data: 2, want: 2},
		{data: 3, want: 6},
		{data: 4, want: 24},
		{data: 5, want: 120},
		{data: -1, want: 0},
	}
	for _, tt := range tests {
		res := FactorialIterative(tt.data)
		if res != tt.want {
			t.Errorf("FactorialIterative(%d)=%d, want %d", tt.data, res, tt.want)
		}
	}
}

func TestFactorialRecursive(t *testing.T) {
	tests := []TestCase{
		{data: 1, want: 1},
		{data: 2, want: 2},
		{data: 3, want: 6},
		{data: 4, want: 24},
		{data: 5, want: 120},
		{data: -1, want: 0},
	}
	for _, tt := range tests {
		res := FactorialRecursive(tt.data)
		if res != tt.want {
			t.Errorf("FactorialIterative(%d)=%d, want %d", tt.data, res, tt.want)
		}
	}
}
