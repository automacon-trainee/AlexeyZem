package list

import (
	"hash"

	"hashMap/internal"
)

type List struct {
	head *Node
	tail *Node
	cap  int
	len  int
}

type Node struct {
	data internal.Data
	next *Node
}

func (l *List) Length() int {
	return l.len
}

func (l *List) Capacity() int {
	return l.cap
}

func (l *List) Rehash(hasher *hash.Hash) internal.Container {
	if l.Capacity() == 0 {
		newList := CreateNewList(3)
		newList.cap = 3
		return newList
	}
	newList := CreateNewList(2 * l.Capacity())
	cur := l.head
	nilData := internal.Data{}
	for i := 0; i < l.Capacity(); i++ {
		if cur.data != nilData {
			hashVal := internal.Hash(cur.data.Key, hasher) % (l.Capacity() * 2)
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
	newList.len = l.Length()
	return newList
}

func (l *List) Set(key string, value any, hasher *hash.Hash) {
	hashVal := internal.Hash(key, hasher)
	hashVal %= l.Capacity()
	cur := l.head
	for i := 0; i < hashVal; i++ {
		cur = cur.next
	}
	nilData := internal.Data{}
	for cur.data != nilData && cur != l.tail && cur.data.Key != key {
		cur = cur.next
	}
	if cur.data != nilData && cur.data.Key != key {
		cur = l.head
		for cur.data != nilData {
			cur = cur.next
		}
	}
	cur.data = internal.Data{Key: key, Value: value}
	l.len++
}

func (l *List) Get(key string, hasher *hash.Hash) (any, bool) {
	hashVal := internal.Hash(key, hasher)
	hashVal %= l.Capacity()
	cur := l.head
	for i := 0; i < hashVal; i++ {
		cur = cur.next
	}
	start := cur
	for cur.data.Key != key && cur != l.tail {
		cur = cur.next
	}
	if cur.data.Key == key {
		return cur.data.Value, true
	} else {
		cur = l.head
		for cur.data.Key != key && cur != start {
			cur = cur.next
		}
		if cur.data.Key == key {
			return cur.data.Value, true
		}
	}
	return nil, false
}

func CreateNewList(len int) *List {
	if len <= 0 {
		return &List{head: nil, tail: nil, cap: 0, len: 0}
	}
	res := List{head: &Node{data: internal.Data{}, next: nil}, tail: &Node{data: internal.Data{}, next: nil}, cap: len}
	cur := res.head
	for i := 1; i < len; i++ {
		cur.next = &Node{data: internal.Data{}, next: nil}
		cur = cur.next
	}
	res.tail = cur
	return &res
}
