package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Node struct {
	data *Commit
	next *Node
	prev *Node
}

type DoubleLinkedList struct {
	head   *Node
	tail   *Node
	curr   *Node
	length int
}

func (d *DoubleLinkedList) Init(c []Commit) {
	d.length = len(c)
	if len(c) == 0 {
		d.head, d.curr, d.tail = nil, nil, nil
		return
	}
	head := &Node{data: &c[0]}
	d.head = head
	d.curr = head
	curr := d.head
	for i := 1; i < len(c); i++ {
		node := &Node{data: &c[i]}
		curr.next = node
		node.prev = curr
		curr = node
	}
	d.tail = curr
}

func (d *DoubleLinkedList) LoadData(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var data []Commit
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}
	QuickSort(data)
	d.Init(data)
	return nil
}

func (d *DoubleLinkedList) Len() int {
	return d.length
}

func (d *DoubleLinkedList) SetCurrent(n int) error {
	if n > d.length || n < 0 {
		return fmt.Errorf("index out of range")
	}
	curr := d.head
	for i := 0; i < n; i++ {
		curr = curr.next
	}
	d.curr = curr
	return nil
}

func (d *DoubleLinkedList) Current() *Node { return d.curr }

func (d *DoubleLinkedList) Next() *Node {
	if d.curr != nil {
		return d.curr.next
	}
	return nil
}

func (d *DoubleLinkedList) Prev() *Node {
	if d.curr != nil {
		return d.curr.prev
	}
	return nil
}

func (d *DoubleLinkedList) Insert(n int, c Commit) error {
	if n > d.length || n < 0 {
		return fmt.Errorf("index out of range")
	}
	newNode := &Node{data: &c}
	if n == 0 {
		if d.head == nil {
			d.head = newNode
			d.tail = newNode
		} else {
			newNode.next = d.head
			d.head.prev = newNode
			d.head = newNode
		}
	} else if n == d.Len() {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	} else {
		curr := d.head
		for i := 0; i < n; i++ {
			curr = curr.next
		}
		newNode.next = curr
		newNode.prev = curr.prev
		curr.prev.next = newNode
		curr.prev = newNode
	}
	d.length++
	return nil
}

func (d *DoubleLinkedList) Push(c Commit) {
	node := &Node{data: &c}
	if d.head == nil {
		d.head = node
		d.curr = node
		d.tail = node
	} else {
		d.tail.next = node
		node.prev = d.tail
		d.tail = node
	}
	d.length++
}

func (d *DoubleLinkedList) Delete(n int) error {
	if n >= d.length || n < 0 {
		return fmt.Errorf("index out of range")
	}
	if n == 0 {
		d.head = d.head.next
		if d.head != nil {
			d.head.prev = nil
		}
	} else if n == d.length-1 {
		d.tail = d.tail.prev
		d.tail.next = nil
	} else {
		curr := d.head
		for i := 0; i < n; i++ {
			curr = curr.next
		}
		if d.curr == curr {
			d.curr = curr.prev
		}
		curr.prev.next = curr.next
		curr.next.prev = curr.prev
	}
	d.length--
	return nil
}

func (d *DoubleLinkedList) DeleteCurrent() error {
	if d.head == nil {
		return fmt.Errorf("head is nil")
	}
	if d.head == d.curr {
		d.head = d.head.next
		if d.head != nil {
			d.head.prev = nil
		}
		d.curr = d.head
	} else if d.tail == d.curr {
		d.tail = d.tail.prev
		d.tail.next = nil
		d.curr = d.tail
	} else {
		d.curr.prev.next = d.curr.next
		d.curr.next.prev = d.curr.prev
		d.curr = d.curr.prev
	}
	d.length--
	return nil
}

func (d *DoubleLinkedList) Index() (int, error) {
	if d.head == nil {
		return 0, fmt.Errorf("head is nil")
	}
	curr := d.head
	var i int
	for i = 0; curr != d.curr; i++ {
		curr = curr.next
	}
	return i, nil
}

func (d *DoubleLinkedList) GetByIndex(n int) (*Node, error) {
	if n >= d.length || n < 0 {
		return nil, fmt.Errorf("index out of range")
	}
	curr := d.head
	for i := 0; i < n; i++ {
		curr = curr.next
	}
	return curr, nil
}

func (d *DoubleLinkedList) Pop() *Node {
	return d.tail
}

func (d *DoubleLinkedList) Shift() *Node {
	return d.head
}

func (d *DoubleLinkedList) SearchUUID(uuid string) *Node {
	curr := d.head
	for curr != nil {
		if curr.data.UUID == uuid {
			return curr
		}
		curr = curr.next
	}
	return nil
}

func (d *DoubleLinkedList) Search(msg string) *Node {
	curr := d.head
	for curr != nil {
		if curr.data.Message == msg {
			return curr
		}
		curr = curr.next
	}
	return nil
}

func (d *DoubleLinkedList) Reverse() *DoubleLinkedList {
	if d.length == 0 {
		return &DoubleLinkedList{}
	}
	cur := d.head
	newHead := &Node{data: d.tail.data}
	newTail := &Node{data: d.head.data}
	newCurr := newTail

	for cur != d.tail {
		newCurr.prev = &Node{data: cur.next.data}
		cur = cur.next
		newCurr.prev.next = newCurr
		newCurr = newCurr.prev
	}

	return &DoubleLinkedList{
		head:   newHead,
		tail:   newTail,
		curr:   d.curr,
		length: d.length,
	}
}

type LinkedLister interface {
	LoadData(path string) error
	Init(c []Commit)
	Len() int
	SetCurrent(n int) error
	Current() *Node
	Next() *Node
	Prev() *Node
	Insert(n int, c Commit) error
	Push(c Commit)
	Delete(n int) error
	DeleteCurrent() error
	Index() (int, error)
	GetByIndex(n int) (*Node, error)
	Pop() *Node
	Shift() *Node
	SearchUUID(uuid string) *Node
	Search(message string) *Node
	Reverse() *DoubleLinkedList
}

type Commit struct {
	UUID    string `json:"uuid"`
	Message string `json:"message"`
	Date    string `json:"date"`
}

func QuickSort(commits []Commit) {
	quickSort(commits, 0, len(commits)-1)
}

func quickSort(list []Commit, low, high int) {
	if low < high {
		pi := partition(list, low, high)
		quickSort(list, low, pi-1)
		quickSort(list, pi+1, high)
	}
}

func partition(list []Commit, low, high int) int {
	pivot := list[high]
	i := low - 1
	for j := low; j < high; j++ {
		if list[j].Date < pivot.Date {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[high] = list[high], list[i+1]
	return i + 1
}

func main() {}
