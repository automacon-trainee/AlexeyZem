package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	text string
	want string
}

func TestCreateUniqueText(t *testing.T) {
	tests := []TestCase{
		{"a b c b c c", "a b c"},
		{"Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum", "Lorem ipsum dolor sit amet consectetur adipiscing elit"},
	}
	for _, test := range tests {
		res := createUniqueText(test.text)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("TestCountWordOccurrences(%s)=%v, want %v", test.text, res, test.want)
		}
	}
}
