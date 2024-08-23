package internal

import (
	"hash"
)

type List struct {
	head *Node
	tail *Node
	cap  int
	len  int
}

type Node struct {
	data Data
	next *Node
}

func (l *List) length() int {
	return l.len
}

func (l *List) capacity() int {
	return l.cap
}

func (l *List) rehash(hasher *hash.Hash) Container {
	if l.capacity() == 0 {
		newList := CreateNewList(3)
		newList.cap = 3
		return newList
	}
	newList := CreateNewList(2 * l.capacity())
	cur := l.head
	nilData := Data{}
	for i := 0; i < l.capacity(); i++ {
		if cur.data != nilData {
			hashVal := Hash(cur.data.key, hasher) % (l.capacity() * 2)
			newCurr := newList.head
			for j := 0; j < hashVal; j++ {
				newCurr = newCurr.next
			}
			for newCurr.data != nilData && newCurr != newList.tail {
				newCurr = newCurr.next
			}
			if newCurr.data != nilData {
				newCurr = newList.head
				for newCurr.data != nilData {
					newCurr = newCurr.next
				}
			}
			newCurr.data = cur.data
		}
		cur = cur.next
	}
	newList.len = l.length()
	return newList
}

func (l *List) Set(key string, value any, hasher *hash.Hash) {
	hashVal := Hash(key, hasher)
	hashVal %= l.capacity()
	cur := l.head
	for i := 0; i < hashVal; i++ {
		cur = cur.next
	}
	nilData := Data{}
	for cur.data != nilData && cur != l.tail && cur.data.key != key {
		cur = cur.next
	}
	if cur.data != nilData && cur.data.key != key {
		cur = l.head
		for cur.data != nilData {
			cur = cur.next
		}
	}
	cur.data = Data{key: key, value: value}
	l.len++
}

func (l *List) Get(key string, hasher *hash.Hash) (any, bool) {
	hashVal := Hash(key, hasher)
	hashVal %= l.capacity()
	cur := l.head
	for i := 0; i < hashVal; i++ {
		cur = cur.next
	}
	start := cur
	for cur.data.key != key && cur != l.tail {
		cur = cur.next
	}
	if cur.data.key == key {
		return cur.data.value, true
	} else {
		cur = l.head
		for cur.data.key != key && cur != start {
			cur = cur.next
		}
		if cur.data.key == key {
			return cur.data.value, true
		}
	}
	return nil, false
}

func CreateNewList(len int) *List {
	if len <= 0 {
		return &List{head: nil, tail: nil, cap: 0, len: 0}
	}
	res := List{head: &Node{data: Data{}, next: nil}, tail: &Node{data: Data{}, next: nil}, cap: len}
	cur := res.head
	for i := 1; i < len; i++ {
		cur.next = &Node{data: Data{}, next: nil}
		cur = cur.next
	}
	res.tail = cur
	return &res
}
