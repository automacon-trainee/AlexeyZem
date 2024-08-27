package main

import (
	"net"

	"chat/internal"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	internal.StartClient(conn)
}
