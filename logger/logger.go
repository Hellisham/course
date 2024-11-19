package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (logger ConsoleLogger) Log(message string) {
	fmt.Println("Your message is: ", message)
}

type FileLogger struct {
	Filepath string
}

func (logger FileLogger) Log(message string) {
	file, err := os.OpenFile(logger.Filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	fmt.Fprintln(file, "Your Message Is :", message)
}
