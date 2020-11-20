package common

import "fmt"

type Logger interface {
	WriteLog()
}

type FileLogger struct {
}

func (*FileLogger) WriteLog() {
	fmt.Println("This is a file log!")
}

type DatabaseLogger struct {
}

func (*DatabaseLogger) WriteLog() {
	fmt.Println("This is a database log!")
}

type LoggerFactory interface {
	CreateLogger() Logger
}

type FileLoggerFactory struct {
}

func (*FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}
