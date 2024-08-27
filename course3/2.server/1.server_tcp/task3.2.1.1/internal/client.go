package internal

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func clientReader(conn net.Conn) {
	for {
		capacityBuff := 1024
		buf := make([]byte, capacityBuff)
		res := ""
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("read error: %v\n", err)
		}
		for n >= capacityBuff {
			res += string(buf[:n])
			n, err = conn.Read(buf)
			if err != nil {
				log.Printf("read error: %v\n", err)
			}
		}
		res += string(buf[:n])
		fmt.Println(res)
	}

}

func StartClient(conn net.Conn) {
	go clientReader(conn)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		_, err := conn.Write([]byte(line + "\r\n"))
		if err != nil {
			fmt.Println(err)
		}
	}
}
