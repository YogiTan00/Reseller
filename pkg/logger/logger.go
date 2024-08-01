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

const (
	Dat      = "[DATA]"
	Err      = "[ERROR]"
	Inf      = "[INFO]"
	Fat      = "[FATAL]"
	Trx      = "[TRX]"
	EndPoint = "[END POINT]"
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
	l := &Logger{
		EndPoint: endPoint,
		log:      log.New(os.Stdout, utils.Color(utils.Green, "[LOGGER] "), log.LstdFlags),
	}
	l.log.Println(utils.Color(utils.Green, EndPoint), endPoint)
	return l
}

func (l *Logger) Info(message any) {
	if l.StatusCode != 0 {
		l.log.Println(utils.Color(utils.Blue, Inf), l.ResponseData)
	} else {
		l.log.Println(utils.Color(utils.Blue, Inf), message)
	}
}

func (l *Logger) Error(err error) {
	if l.isError {
		l.log.Println(utils.Color(utils.Red, Err), l.ResponseData)
	} else {
		l.log.Println(utils.Color(utils.Red, Err), err)
	}
}

func (l *Logger) InfoWithData(message string) {
	l.log.Println(utils.Color(utils.Blue, Inf), strings.Title(message))
	if l.ResponseData != nil {
		l.log.Println(utils.Color(utils.Yellow, Dat), l.ResponseData)
	}
}

func (l *Logger) CreateNewLog() {
	var now = time.Now()
	if _, err := os.Stat(fmt.Sprintf("%s/logs", os.TempDir())); os.IsNotExist(err) {
		if err = os.MkdirAll(fmt.Sprintf("%s/logs", os.TempDir()), 0755); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}
	file, err := os.OpenFile(fmt.Sprintf("%s/logs/%s.log", os.TempDir(), time.Now().Format(time.DateOnly)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l.Info(err)
	}
	defer file.Close()
	l.log = log.New(file, now.Format(time.DateTime), log.LstdFlags)
	if !l.isError && l.StatusCode == 0 {
		l.StatusCode = http.StatusOK
	}
	if l.TrxId == "" {
		l.TrxId = "TRX_ID_IS_EMPTY"
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

func (l *Logger) Fatal(err error) {
	if l.isError {
		l.log.Println(utils.Color(utils.Red, Fat), l.ResponseData)
	} else {
		l.log.Println(utils.Color(utils.Red, Fat), err)
	}
	os.Exit(1)
}
