package main

import (
	"github.com/google/btree"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func (u User) Less(than btree.Item) bool {
	return u.ID < than.(User).ID
}

type BTree struct {
	tree *btree.BTree
}

func NewBTree(degree int) *BTree {
	return &BTree{
		tree: btree.New(degree),
	}
}

func (bt *BTree) Insert(user User) {
	bt.tree.ReplaceOrInsert(user)
}

func (bt *BTree) Search(id int) *User {
	i := bt.tree.Get(User{ID: id})
	if i == nil {
		return nil
	}
	return &User{ID: i.(User).ID, Name: i.(User).Name, Age: i.(User).Age}
}

func main() {}
