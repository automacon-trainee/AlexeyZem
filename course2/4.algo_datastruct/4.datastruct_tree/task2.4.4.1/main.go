package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type Node struct {
	left  *Node
	right *Node
	data  *User
}

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) Insert(data *User) {
	if t.root == nil {
		t.root = &Node{data: data}
		return
	}
	t.root.Insert(data)
}

func (n *Node) Insert(data *User) {
	if data.ID > n.data.ID {
		if n.right != nil {
			n.right.Insert(data)
		} else {
			n.right = &Node{data: data}
		}
	} else {
		if n.left != nil {
			n.left.Insert(data)
		} else {
			n.left = &Node{data: data}
		}
	}
}

func (t *BinaryTree) Search(key int) *User {
	if t.root == nil {
		return nil
	}
	return t.root.Search(key)
}

func (n *Node) Search(key int) *User {
	if n == nil {
		return nil
	}
	if n.data.ID == key {
		return n.data
	}
	if n.data.ID > key {
		return n.left.Search(key)
	} else {
		return n.right.Search(key)
	}
}

func GenerateData(n int) *BinaryTree {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	bt := &BinaryTree{}
	for i := 0; i < n; i++ {
		val := rand.Intn(100)
		bt.Insert(&User{
			ID:   val,
			Name: fmt.Sprintf("User%d", val),
			Age:  rand.Intn(50) + 20,
		})
	}
	return bt
}

func PrintTree(bt *BinaryTree) {
	PrintNode(bt.root)
}

func PrintNode(n *Node) {
	if n.left != nil {
		PrintNode(n.left)
	}
	fmt.Println(n.data)
	if n.right != nil {
		PrintNode(n.right)
	}
}

func main() {
	count := 30
	bt := GenerateData(count)
	user := bt.Search(count)
	if user != nil {
		fmt.Println(user)
	} else {
		fmt.Println("User not found")
	}
	PrintTree(bt)
}
