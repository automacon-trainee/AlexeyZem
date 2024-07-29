package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	data []User
	want []User
}

func TestGetUniqueUser(t *testing.T) {
	tests := []TestCase{
		{
			[]User{{"Aba", 19, ""}, {"Aba", 19, ""}, {"Ars", 20, ""}},
			[]User{{"Aba", 19, ""}, {"Ars", 20, ""}},
		},
	}
	for _, test := range tests {
		res := getUniqueUser(test.data)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("getUniqueUser(%v) = %v; want %v", test.data, res, test.want)
		}
	}
}
