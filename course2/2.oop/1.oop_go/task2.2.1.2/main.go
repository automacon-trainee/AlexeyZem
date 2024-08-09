package main

import (
	"fmt"
	"io"
	"os"
)

type Logger interface {
	Log(msg string) error
}

type FileLogger struct {
	File *os.File
}

func (fl *FileLogger) Log(msg string) error {
	_, err := fl.File.WriteString(msg)
	return err
}

type ConsoleLogger struct {
	Writer io.Writer
}

func (cl *ConsoleLogger) Log(msg string) error {
	_, err := cl.Writer.Write([]byte(msg))
	return err
}

func LogAll(loggers []Logger, msg string) {
	for _, logger := range loggers {
		err := logger.Log(msg)
		if err != nil {
			fmt.Println("Failed to log message:", err)
		}
	}
}

func main() {
	consoleLogger := &ConsoleLogger{Writer: os.Stdout}
	file, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}
	fileLogger := &FileLogger{File: file}
	LogAll([]Logger{consoleLogger, fileLogger}, "test log message\n")
}
