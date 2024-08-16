package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func writeFile(data io.Reader, fd io.Writer) error {
	length := 1024
	buf := make([]byte, length)
	for n, err := data.Read(buf); !errors.Is(err, io.EOF); n, err = data.Read(buf) {
		buf = buf[:n]
		_, err = fd.Write(buf)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	filePath := "aa.txt"
	numPerm := 0666
	perm := os.FileMode(numPerm)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		fmt.Println("open file err:", err)
	}
	defer file.Close()
	err = writeFile(strings.NewReader("Hello World\nПривет мир\nGO the best"), file)
	if err != nil {
		fmt.Println("write file err:", err)
	}
}
