package main

import (
	"reflect"
	"testing"
)

type TestFilter struct {
	users []User
	want  []User
}

func TestFilterComments(t *testing.T) {
	usersBad := []User{
		{
			Name: "Betty",
			Comments: []Comment{
				{Message: "good Comment 1"},
				{Message: "BaD CoMmEnT 2"},
				{Message: "Bad Comment 3"},
				{Message: "Use camelCase please"},
			},
		},
		{
			Name: "Jhon",
			Comments: []Comment{
				{Message: "Good Comment 1"},
				{Message: "Good Comment 2"},
				{Message: "Good Comment 3"},
				{Message: "Bad Comments 4"},
			},
		},
	}
	users := []User{
		{
			Name: "Betty",
			Comments: []Comment{
				{Message: "good Comment 1"},
				{Message: "BaD CoMmEnT 2"},
				{Message: "Bad Comment 3"},
				{Message: "Use camelCase please"},
			},
		},
		{
			Name: "Jhon",
			Comments: []Comment{
				{Message: "Good Comment 1"},
			},
		},
	}
	tests := []TestFilter{
		{usersBad, []User{}},
		{users, []User{{Name: "Jhon", Comments: []Comment{{Message: "Good Comment 1"}}}}},
	}
	for _, test := range tests {
		res := FilterComments(test.users)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("FilterComments(%v) got %v want %v", test.users, res, test.want)
		}
	}
}

type TestCaseIsBad struct {
	comment Comment
	want    bool
}

func TestIsBadComment(t *testing.T) {
	tests := []TestCaseIsBad{
		{Comment{Message: "good Comment 1"}, false},
		{Comment{Message: "bad coMMenT 2"}, true},
	}
	for _, test := range tests {
		res := IsBadComment(test.comment)
		if res != test.want {
			t.Errorf("IsBadComment(%v) got %v want %v", test.comment, res, test.want)
		}
	}
}

type TestCaseGetBad struct {
	users []User
	want  []Comment
}

func TestGetBadComment(t *testing.T) {
	usersBad := []User{
		{
			Name: "Betty",
			Comments: []Comment{
				{Message: "good Comment 1"},
				{Message: "BaD CoMmEnT 2"},
				{Message: "Bad Comment 3"},
				{Message: "Use camelCase please"},
			},
		},
		{
			Name: "Jhon",
			Comments: []Comment{
				{Message: "Good Comment 1"},
				{Message: "Good Comment 2"},
				{Message: "Good Comment 3"},
				{Message: "Bad Comments 4"},
			},
		},
	}
	tests := []TestCaseGetBad{
		{users: usersBad, want: []Comment{
			{Message: "BaD CoMmEnT 2"},
			{Message: "Bad Comment 3"},
			{Message: "Bad Comments 4"},
		},
		},
	}
	for _, test := range tests {
		res := GetBadComments(test.users)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("GetBadComments(%v) got %v want %v", test.users, res, test.want)
		}
	}
}

type TestCaseGetGood struct {
	users []User
	want  []Comment
}

func TestGetGoodComment(t *testing.T) {
	users := []User{
		{
			Name: "Betty",
			Comments: []Comment{
				{Message: "good Comment 1"},
				{Message: "BaD CoMmEnT 2"},
				{Message: "Bad Comment 3"},
				{Message: "Use camelCase please"},
			},
		},
		{
			Name: "Jhon",
			Comments: []Comment{
				{Message: "Good Comment 1"},
			},
		},
	}
	tests := []TestCaseGetGood{
		{users: users, want: []Comment{
			{Message: "good Comment 1"},
			{Message: "Use camelCase please"},
			{Message: "Good Comment 1"},
		}},
	}
	for _, test := range tests {
		res := GetGoodComment(test.users)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("GetGoodComment(%v) got %v want %v", test.users, res, test.want)
		}
	}
}
