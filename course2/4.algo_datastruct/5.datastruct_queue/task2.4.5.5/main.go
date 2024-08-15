package main

import (
	"fmt"
)

type CircuitRinger interface {
	Add(value int)
	Get() (int, bool)
}

type Buff struct {
	capacity int
	data     []int
	length   int
	cur      int
}

func (b *Buff) Add(value int) {
	if b.length < b.capacity {
		b.data = append(b.data, value)
		b.length++
	} else {
		b.data[b.cur] = value
	}
	b.cur = (b.cur + 1) % b.capacity
}
func (b *Buff) Get() (int, bool) {
	if b.length == 0 {
		return 0, false
	}
	if b.cur >= b.length {
		b.cur = 0
	}
	result := b.data[b.cur]
	b.data = append(b.data[0:b.cur], b.data[b.cur+1:]...)
	b.length--
	return result, true
}

func NewRingBuffer(capacity int) *Buff {
	return &Buff{capacity: capacity, data: make([]int, 0, capacity), length: 0, cur: 0}
}

func main() {
	rb := NewRingBuffer(3)
	rb.Add(1)
	rb.Add(2)
	rb.Add(3)
	rb.Add(4)
	for val, ok := rb.Get(); ok; val, ok = rb.Get() {
		fmt.Println(val)
	}
}
