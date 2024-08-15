package main

import (
	"testing"
)

func TestBTree(t *testing.T) {
	bt := NewBTree(3)
	users := []User{
		{1, "Alice", 30},
		{2, "Bob", 20},
		{3, "Charlie", 40},
	}
	for _, user := range users {
		bt.Insert(user)
	}
	user := bt.Search(1)
	if *user != users[0] {
		t.Errorf("search fail")
	}
	user = bt.Search(10)
	if user != nil {
		t.Errorf("search fail")
	}
}
