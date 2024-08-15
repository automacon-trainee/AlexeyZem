package main

import (
	"os"
)

type Logger interface {
	Log(msg string) error
}

type FileLogger struct {
	file *os.File
}

func (l FileLogger) Log(msg string) error {
	_, err := l.file.WriteString(msg)
	if err != nil {
		return err
	}
	return nil
}

type LogSystem struct {
	logger Logger
}

type LogOptions func(*LogSystem)

func WithLogger(logger Logger) LogOptions {
	return func(l *LogSystem) {
		l.logger = logger
	}
}
func NewLogSystem(opts ...LogOptions) *LogSystem {
	ls := &LogSystem{}
	for _, opt := range opts {
		opt(ls)
	}
	return ls
}

func main() {}
