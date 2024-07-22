package logger

import (
	"fmt"
	"github.com/YogiTan00/Reseller/pkg/utils"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Logger struct {
	isError      bool
	log          *log.Logger
	TimeStarted  time.Time
	EndPoint     string
	StatusCode   int
	RequestData  interface{}
	ResponseData interface{}
	TrxId        string
}

func NewLogger(endPoint string) *Logger {
	return &Logger{
		EndPoint: endPoint,
		log:      log.New(os.Stdout, "[Logger] ", log.LstdFlags),
	}
}

func (l *Logger) Info(message any) {
	if l.StatusCode != 0 {
		l.log.Println(utils.Color("green", "[EndPoint]"), l.EndPoint)
		l.log.Println(utils.Color("blue", "[INFO]"), l.ResponseData)
	} else {
		l.log.Println(utils.Color("green", "[EndPoint]"), l.EndPoint)
		l.log.Println(utils.Color("blue", "[INFO]"), message)
	}
}

func (l *Logger) Error(err error) {
	if l.isError {
		l.log.Println(utils.Color("green", "[EndPoint]"), l.EndPoint)
		l.log.Println(utils.Color("red", "[ERROR]"), l.ResponseData)
	} else {
		l.log.Println(utils.Color("green", "[EndPoint]"), l.EndPoint)
		l.log.Println(utils.Color("red", "[ERROR]"), err)
	}
}

func (l *Logger) InfoWithData(message string) {
	if l.EndPoint != "" {
		l.log.Println(utils.Color("green", "[EndPoint]"), l.EndPoint)
	}
	l.log.Println(utils.Color("blue", "[INFO]"), strings.Title(message))
	if l.ResponseData != nil {
		l.log.Println(utils.Color("yellow", "[Data]"), l.ResponseData)
	}
}

func (l *Logger) CreateNewLog() {
	var now = time.Now()
	file, err := os.OpenFile(fmt.Sprintf("%s", now.Format(time.DateOnly)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()
	l.log = log.New(file, now.Format(time.DateTime), log.LstdFlags)
	if !l.isError && l.StatusCode == 0 {
		l.StatusCode = http.StatusOK
	}
	if l.TrxId == "" {
		l.TrxId = "trx_id_is_empty"
	}

	if l.StatusCode >= 500 && l.StatusCode <= 599 {
		l.MarkAsError()
	}

	if l.isError {
		l.Error(nil)
	} else {
		l.Info(nil)
	}

}

func (l *Logger) MarkAsError() {
	l.isError = true
}
