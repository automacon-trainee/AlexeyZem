package main

import (
	"fmt"
)

type Bank struct {
	queue []string
}

func (b *Bank) AddClient(client string) {
	b.queue = append(b.queue, client)
}

func (b *Bank) ServeNextClient() string {
	if len(b.queue) != 0 {
		res := b.queue[0]
		b.queue = b.queue[1:]
		return res
	}
	return ""
}

func main() {
	b := &Bank{}
	c := b.ServeNextClient()
	fmt.Println(c)
	b.AddClient("Alexey")
	b.AddClient("Andrew")
	b.AddClient("Author")
	fmt.Println(b.ServeNextClient())
	fmt.Println(b.ServeNextClient())
	fmt.Println(b.ServeNextClient())
}
