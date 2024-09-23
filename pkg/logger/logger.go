package logger

import (
	"encoding/json"
	"fmt"
	"github.com/YogiTan00/Reseller/config"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"strconv"
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
	prefix       string
	log          *logrus.Logger
	EndPoint     string
	StatusCode   int
	RequestData  interface{}
	ResponseData interface{}
	TrxId        string
}

func (l *Logger) Info(message any) {
	if l.EndPoint != "" {
		l.log.Println(l.EndPoint)
	}
	if l.StatusCode != 0 {
		if l.ResponseData != nil {
			result, _ := json.Marshal(l.ResponseData)
			l.log.Println(l.TrxId, string(result))
		}
	} else {
		l.log.Println(message)
	}
}

func (l *Logger) Error(err error) {
	if l.isError {
		l.log.Errorln(l.ResponseData)
	} else {
		l.log.Errorln(err)
	}
}

func (l *Logger) InfoWithData(message string, data any) {
	l.log.Println(strings.Title(message))
	if data != nil {
		l.log.Println(data)
	}
	if l.ResponseData != nil {
		l.log.Println(l.ResponseData)
	}
}

func (l *Logger) CreateNewLog() {
	var (
		debug bool
		err   error
		cfg   = config.NewConfig()
	)
	l.log = logrus.New()
	l.log.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}
	if len(cfg.Debug) != 0 {
		debug, err = strconv.ParseBool(cfg.Debug)
		if err != nil {
			l.Fatal(err)
		}
	}
	if debug {
		saveLog(l, cfg)
	}
	createLog(l)
}

func createLog(l *Logger) {
	l.log.Out = os.Stdout
	l.log.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}
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

func saveLog(l *Logger, cfg *config.Config) {
	if _, err := os.Stat(fmt.Sprintf("%s/logs", cfg.PathLogs)); os.IsNotExist(err) {
		if err = os.MkdirAll(fmt.Sprintf("%s/logs", cfg.PathLogs), 0755); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}
	file, err := os.OpenFile(fmt.Sprintf("%s/logs/%s.log", cfg.PathLogs, time.Now().Format(time.DateOnly)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l.Info(err)
	}
	defer file.Close()
	l.log.Out = file
	l.log.Formatter = &logrus.TextFormatter{
		ForceColors:   false,
		FullTimestamp: true,
	}
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
		l.log.Println(Fat, l.ResponseData)
	} else {
		l.log.Println(Fat, err)
	}
	os.Exit(1)
}

func (l *Logger) Panic(err error) {
	if l.isError {
		l.log.Println(Fat, l.ResponseData)
	} else {
		l.log.Println(Fat, err)
	}
	panic(err)
}
