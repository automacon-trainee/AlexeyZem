package internal

import (
	"encoding/binary"
	"hash"
	"hash/crc32"
	"hash/crc64"
)

type HashMapList struct {
	hasher   hash.Hash
	list     *List
	length   int
	capacity int
}

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	data Data
	next *Node
}

func (h *HashMapList) rehash() {
	h.hasher.Reset()
	if h.capacity == 0 {
		h.list = CreateNewList(3)
		h.capacity = 3
		return
	}
	newList := CreateNewList(2 * h.capacity)
	cur := h.list.head
	nilData := Data{}
	for i := 0; i < h.capacity; i++ {
		if cur.data != nilData {
			hashVal := h.Hash(cur.data.key) % (h.capacity * 2)
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
	h.capacity *= 2
	h.list = newList
}

func (h *HashMapList) Set(key string, value any) {
	hashVal := h.Hash(key)
	if 3*h.capacity <= h.length*4 {
		h.rehash()
	}
	hashVal %= h.capacity
	cur := h.list.head
	for i := 0; i < hashVal; i++ {
		cur = cur.next
	}
	nilData := Data{}
	for cur.data != nilData && cur != h.list.tail && cur.data.key != key {
		cur = cur.next
	}
	if cur.data != nilData && cur.data.key != key {
		cur = h.list.head
		for cur.data != nilData {
			cur = cur.next
		}
	}
	cur.data = Data{key: key, value: value}
	h.length++
}

func (h *HashMapList) Get(key string) (any, bool) {
	hashVal := h.Hash(key)
	if 3*h.capacity <= h.length*4 {
		h.rehash()
	}
	hashVal %= h.capacity
	cur := h.list.head
	for i := 0; i < hashVal; i++ {
		cur = cur.next
	}
	start := cur
	for cur.data.key != key && cur != h.list.tail {
		cur = cur.next
	}
	if cur.data.key == key {
		return cur.data.value, true
	} else {
		cur = h.list.head
		for cur.data.key != key && cur != start {
			cur = cur.next
		}
		if cur.data.key == key {
			return cur.data.value, true
		}
	}
	return nil, false
}

func (h *HashMapList) Hash(key string) int {
	h.hasher.Reset()
	h.hasher.Write([]byte(key))
	return int(binary.BigEndian.Uint16(h.hasher.Sum(nil)))
}

func NewHashMapList(capacity int, opt OptionList) *HashMapList {
	list := CreateNewList(capacity)
	res := &HashMapList{list: list, length: 0, capacity: capacity}
	opt(res)
	return res
}

func CreateNewList(len int) *List {
	if len <= 0 {
		return nil
	}
	res := List{&Node{data: Data{}, next: nil}, &Node{data: Data{}, next: nil}}
	cur := res.head
	for i := 1; i < len; i++ {
		cur.next = &Node{data: Data{}, next: nil}
		cur = cur.next
	}
	res.tail = cur
	return &res
}

type OptionList func(slice *HashMapList)

func WithHashCRC64List() OptionList {
	return func(h *HashMapList) {
		h.hasher = crc64.New(crc64.MakeTable(crc64.ECMA))
	}
}

func WithHashCRC32List() OptionList {
	return func(h *HashMapList) {
		h.hasher = crc32.NewIEEE()
	}
}
