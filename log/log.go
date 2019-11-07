package common

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

type SelfLog struct {
	log.Logger
}

func init() {
	f, err := NewRollingFile("fool_dealer_", "2006-01-02", ".log")
	if err != nil {
		fmt.Println("error create file")
	}
	Warning = log.New(io.MultiWriter(f, os.Stderr), "warning ", log.Ldate|log.Ltime|log.Lshortfile)
}
