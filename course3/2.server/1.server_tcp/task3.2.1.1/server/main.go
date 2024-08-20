// У тебя есть код сервера и клиента для простого чата на TCP.
// Задача — добавить недостающие части кода и доработать функциональность чата.
// Критерии приемки
// Сервер должен принимать входящие соединения на порту 8000.
// Когда клиент подключается, сервер должен отправить сообщение «You are [адрес клиента]»
// Когда клиент отправляет сообщение, сервер должен отправить его всем подключенным клиентам в
// формате «[адрес клиента]: [сообщение]».
// Когда клиент отключается, сервер должен отправить сообщение «[адрес клиента] has left».
// Клиент должен подключаться к серверу на порту 8000.
// Клиент должен отправлять введенные с клавиатуры сообщения на сервер.
// Клиент должен выводить полученные сообщения от сервера на экран.
// Для покрытия тестами функционал должен быть выделен в отдельные структуры и методы.
// Функционал должен быть покрыт тестами на 80%

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

type Client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

var (
	entering = make(chan Client)
	leaving  = make(chan Client)
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[Client]bool)

	var mu sync.Mutex
	go func() {
		for enter := range entering {
			mu.Lock()
			clients[enter] = true
			log.Println("entering", enter)
			mu.Unlock()
		}
	}()

	go func() {
		for leav := range leaving {
			mu.Lock()
			delete(clients, leav)
			log.Println("Client", leav, "leaving")
			mu.Unlock()
		}
	}()

	for msg := range messages {
		mu.Lock()
		for client := range clients {
			client.ch <- msg
			log.Println("sending", msg, "to", client)
		}
		mu.Unlock()
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	cli := Client{conn: conn, name: who, ch: ch}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for text := range ch {
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println(err)
		}
	}
}
