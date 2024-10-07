package utils

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{log.New(os.Stdout, "Go-Commerce", log.LstdFlags)}
}

func (l *Logger) Info(message string) {
	l.Printf("[INFO] %s\n", message)
}
