package main

import (
	"strings"
)

type User struct {
	Name     string
	Comments []Comment
}

type Comment struct {
	Message string
}

func FilterComments(users []User) []User {
	res := make([]User, 0)
first:
	for _, user := range users {
		for _, comment := range user.Comments {
			if IsBadComment(comment) {
				continue first
			}
		}
		res = append(res, user)
	}
	return res
}

func IsBadComment(comment Comment) bool {
	return strings.Contains(strings.ToLower(comment.Message), "bad comment")
}

func GetBadComments(users []User) []Comment {
	res := make([]Comment, 0)
	for _, user := range users {
		for _, comment := range user.Comments {
			if IsBadComment(comment) {
				res = append(res, comment)
			}
		}
	}
	return res
}

func GetGoodComment(users []User) []Comment {
	res := make([]Comment, 0)
	for _, user := range users {
		for _, comment := range user.Comments {
			if !IsBadComment(comment) {
				res = append(res, comment)
			}
		}
	}
	return res
}

func main() {}
