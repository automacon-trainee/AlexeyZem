package main

import (
	"os"
	"testing"
)

func TestNewLogSystem(t *testing.T) {
	file, _ := os.Create("log.txt")

	consLog := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(consLog))
	if logSystem.logger != consLog {
		t.Errorf("NewLogSystem wrong")
	}
	str := "hello world"
	err := logSystem.logger.Log(str)
	if err != nil {
		t.Errorf("Logging wrong")
	}
	file.Close()

	file, err = os.OpenFile("log.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Errorf("OpenFile wrong")
	}
	res := make([]byte, 1024)
	n, err := file.Read(res)
	if err != nil {
		t.Errorf("Reading wrong")
	}
	res = res[:n]
	if string(res) != str {
		t.Errorf("wrong result expected: %s, got: %s", str, string(res))
	}
	file.Close()

	err = logSystem.logger.Log(str)
	if err == nil {
		t.Errorf("Logging wrong")
	}
}
