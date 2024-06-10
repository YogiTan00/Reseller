package logger

import (
	"Reseller/pkg/utils"
	"log"
	"os"
	"strings"
)

type Logger struct {
	log      *log.Logger
	endPoint string
}

func NewLogger(endPoint string) *Logger {
	return &Logger{
		endPoint: endPoint,
		log:      log.New(os.Stdout, "[Logger] ", log.LstdFlags),
	}
}

func (l *Logger) Info(message string) {
	l.log.Println(utils.Color("green", "[EndPoint]"), l.endPoint)
	l.log.Println(utils.Color("blue", "[INFO]"), message)
}

func (l *Logger) Error(err error) {
	l.log.Println(utils.Color("green", "[EndPoint]"), l.endPoint)
	l.log.Println(utils.Color("red", "[ERROR]"), err)
}

func (l *Logger) InfoWithData(message string, data any) {
	if l.endPoint != "" {
		l.log.Println(utils.Color("green", "[EndPoint]"), l.endPoint)
	}
	l.log.Println(utils.Color("blue", "[INFO]"), strings.Title(message))
	if data != nil {
		l.log.Println(utils.Color("yellow", "[Data]"), data)
	}
}
