package main

import (
	"fmt"
)

const (
	id1 = 2 * (iota + 1)
	id2
	id3
	id4
	id5
	id6
	id7
	id8
	id9
	id10
	id11
	id12
	id13
)

func main() {
	fmt.Println(id1, id2, id3, id4, id5, id6, id7, id8, id9, id10, id11, id12, id13)
}
