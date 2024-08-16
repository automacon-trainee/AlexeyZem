package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func writeFile(filename string, data []byte, perm os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(filename), perm)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		err2 := f.Close()
		if err2 != nil {
			return fmt.Errorf("%w and %w", err, err2)
		}
		return err
	}
	return f.Close()
}

func main() {
	err := writeFile("test/aa.txt", []byte(" world"), 0777)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
