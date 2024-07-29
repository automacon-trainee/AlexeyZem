package main

import (
	"testing"
)

type TestCase struct {
	operate func(...interface{}) interface{}
	expect  interface{}
	data    []interface{}
}

func TestOperate(t *testing.T) {
	test := []TestCase{
		{Concat, "Hello world", []interface{}{"Hello", " ", "world"}},
		{Sum, 12.0, []interface{}{10.0, 2.0}},
		{Sum, 10, []interface{}{3, 4, 1, 2}},
	}
	for _, v := range test {
		result := Operate(v.operate, v.data...)
		if result != v.expect {
			t.Errorf("Operate false, expect %v, got %v", v.expect, result)
		}
	}
}
