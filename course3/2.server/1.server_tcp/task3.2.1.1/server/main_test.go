package main

import (
	"net"
	"testing"
	"time"
)

func TestBroadcaster(t *testing.T) {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	cl1 := Client{ch: ch1}
	cl2 := Client{ch: ch2}
	go broadcaster()
	entering <- cl1
	entering <- cl2
	messages <- "Hello World"

	if <-ch1 != "Hello World" {
		t.Errorf("ch1 != \"Hello World\"")
	}
	if <-ch2 != "Hello World" {
		t.Errorf("ch2 != \"Hello World\"")
	}
}

type mockConn struct {
	net.Conn
	readData  string
	writeData string
}

func (m *mockConn) Read(b []byte) (n int, err error) {
	copy(b, m.readData)
	return len(m.readData), nil
}

func (m *mockConn) Write(b []byte) (n int, err error) {
	m.writeData = string(b)
	return len(b), nil
}

func (m *mockConn) Close() error {
	return nil
}

func (m *mockConn) RemoteAddr() net.Addr {
	return &net.IPAddr{IP: net.ParseIP("127.0.0.1")}
}

func TestClientWriter(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "Hello World"
	conn := &mockConn{}
	go clientWriter(conn, ch)
	time.Sleep(1 * time.Second)
	str := conn.writeData
	if str != "Hello World" {
		t.Errorf("str = %s, want Hello World", str)
	}
}

func TestHandleConn(t *testing.T) {
	conn := &mockConn{}
	go handleConn(conn)
	time.Sleep(1 * time.Second)
	if conn.writeData == "" {
		t.Errorf("conn.writeData is empty")
	}
}
