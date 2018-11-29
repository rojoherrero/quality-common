package common

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

type Logger interface {
	Error(msg string)
	Info(msg string)
}

type logger struct {
	info *log.Logger
	error *log.Logger
}

func InitLogger(path string) Logger {
	fileName := filepath.Join(path, "/quality-", time.Now().Format("20060102150405"), ".log")
	file, e := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if e != nil {
		panic(e)
	}

	error := log.New(file, "ERROR", log.Ldate|log.Ltime|log.Lshortfile)
	info := log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)

	return &logger{info, error}
}

func (l *logger) Error(msg string){
	l.error.Fatalln(msg)
}

func (l *logger) Info(msg string){
	l.info.Println(msg)
}
