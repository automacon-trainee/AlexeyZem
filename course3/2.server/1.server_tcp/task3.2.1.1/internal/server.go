package internal

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

type Client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

func StartServer(listener net.Listener) {
	log.Println("Start server")
	chat := NewChat()
	go broadcaster(chat)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn, chat)
	}
}

func broadcaster(chat *Chat) {
	clients := make(map[Client]bool)

	var mu sync.Mutex
	go func() {
		for enter := range chat.entering {
			mu.Lock()
			clients[enter] = true
			log.Println("entering", enter)
			mu.Unlock()
		}
	}()

	go func() {
		for leav := range chat.leaving {
			mu.Lock()
			delete(clients, leav)
			log.Println("Client", leav, "leaving")
			mu.Unlock()
		}
	}()

	for msg := range chat.msg {
		mu.Lock()
		for client := range clients {
			client.ch <- msg
			log.Println("sending", msg, "to", client)
		}
		mu.Unlock()
	}
}

func handleConn(conn net.Conn, chat *Chat) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	cli := Client{conn: conn, name: who, ch: ch}
	ch <- "You are " + who
	chat.msg <- who + " has arrived"
	chat.entering <- cli
	input := bufio.NewScanner(conn)
	for input.Scan() {
		chat.msg <- who + ": " + input.Text()
	}
	chat.leaving <- cli
	chat.msg <- who + " has left"
	conn.Close()
	close(ch)
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for text := range ch {
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println(err)
		}
	}
}
