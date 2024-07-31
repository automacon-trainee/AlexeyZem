package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	text  string
	words []string
	want  map[string]int
}

func TestCountWordsInText(t *testing.T) {
	txt := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed
 sit 
amet ipsum mauris. 
        Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et m
ollis dolor. 
        Praesent et diam eget libero egestas mattis sit amet vitae augue.`
	tests := []TestCase{
		{txt, []string{"sit", "amet", "lorem"}, map[string]int{
			"amet": 2, "lorem": 1, "sit": 3,
		}},
		{txt, []string{"donec", "et"}, map[string]int{
			"donec": 2, "et": 2,
		}},
	}
	for _, test := range tests {
		res := CountWordsInText(test.text, test.words)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("CountWordsInText(%s, %v) is %v, want %v", test.text, test.words, res, test.want)
		}
	}
}
