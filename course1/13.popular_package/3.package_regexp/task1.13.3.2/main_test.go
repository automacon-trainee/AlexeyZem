package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	ads    []Ad
	censor map[string]string
	want   []Ad
}

func TestCensorAd(t *testing.T) {
	testCases := []TestCase{
		{
			ads: []Ad{
				{Title: "Куплю велосипед MeRiDa",
					Description: "Куплю велосипед meriDA в хорошем состоянии."},
			}, censor: map[string]string{"велосипед merida": "телефон Apple"},
			want: []Ad{
				{Title: "Куплю телефон Apple",
					Description: "Куплю телефон Apple в хорошем состоянии."},
			},
		},
	}
	for _, tc := range testCases {
		if res := censorAds(tc.ads, tc.censor); !reflect.DeepEqual(res, tc.want) {
			t.Errorf("censorAd wrong. got: %v; want: %v", res, tc.want)
		}
	}
}
