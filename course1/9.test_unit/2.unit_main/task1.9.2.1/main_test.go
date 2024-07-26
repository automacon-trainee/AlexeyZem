package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	err := w.Close()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	os.Stdout = old

	var stdout bytes.Buffer
	_, err = stdout.ReadFrom(r)

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	expected := "Hello, World!\n"
	if stdout.String() != expected {
		t.Errorf("Got %s, expected %s", stdout.String(), expected)
	}
}
