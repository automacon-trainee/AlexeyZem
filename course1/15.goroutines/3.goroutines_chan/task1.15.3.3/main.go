package main

import (
	"fmt"
)

type Sema chan struct{}

func New(n int) Sema {
	return make(Sema, n)

}

func (s Sema) Inc(k int) {
	for i := 0; i < k; i++ {
		s <- struct{}{}
	}
}

func (s Sema) Dec(k int) {
	for i := 0; i < k; i++ {
		<-s
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)
	sem := New(n)
	for _, num := range numbers {
		go func(n int) {
			fmt.Printf("%d ", n)
			sem.Inc(1)
		}(num)
	}
	sem.Dec(n)
}
