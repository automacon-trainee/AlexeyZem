package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	userL []User
	userR []User
	want  []User
}

func TestMerge(t *testing.T) {
	usersL := []User{
		{ID: 1, Name: "skdj", Age: 30},
		{ID: 2, Name: "skdj", Age: 30},
		{ID: 5, Name: "skdj", Age: 30},
		{ID: 10, Name: "skdj", Age: 30},
		{ID: 12, Name: "skdj", Age: 30},
	}
	usersR := []User{
		{ID: 3, Name: "skdj", Age: 30},
		{ID: 4, Name: "skdj", Age: 30},
		{ID: 6, Name: "skdj", Age: 30},
		{ID: 11, Name: "skdj", Age: 30},
		{ID: 13, Name: "skdj", Age: 30},
	}
	users := []User{
		{ID: 1, Name: "skdj", Age: 30},
		{ID: 2, Name: "skdj", Age: 30},
		{ID: 3, Name: "skdj", Age: 30},
		{ID: 4, Name: "skdj", Age: 30},
		{ID: 5, Name: "skdj", Age: 30},
		{ID: 6, Name: "skdj", Age: 30},
		{ID: 10, Name: "skdj", Age: 30},
		{ID: 11, Name: "skdj", Age: 30},
		{ID: 12, Name: "skdj", Age: 30},
		{ID: 13, Name: "skdj", Age: 30},
	}
	tests := []testCase{
		{userL: usersL, userR: usersR, want: users},
		{userL: usersL, userR: []User{}, want: usersL},
		{userL: []User{}, userR: usersR, want: usersR},
		{userL: []User{}, userR: []User{}, want: []User{}},
	}
	for _, tt := range tests {
		res := Merge(tt.userL, tt.userR)
		if !reflect.DeepEqual(res, tt.want) {
			t.Errorf("want: %v, got: %v", tt.want, res)
		}
	}
}
