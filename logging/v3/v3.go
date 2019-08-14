package v3

import (
	"encoding/json"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"log"
	"os"
	"time"
)

type RotateType int

const (
	RotateDefault RotateType = iota
	RotateMinute
	RotateHour
	RotateDay
)

type KV map[string]interface{}

const (
	InfoStr  = "INFO "
	ErrorStr = "ERROR "
)

var logger *JLogger

type JLogger struct {
	rotateType    RotateType
	stdLogger     *log.Logger
	fileLogger    *log.Logger
	enableFileLog bool
}

func init() {
	logger = New()
}

func New() *JLogger {
	stdLogger := log.New(os.Stdout, "", log.LstdFlags)
	return &JLogger{
		rotateType:    RotateDefault,
		stdLogger:     stdLogger,
		enableFileLog: false,
		fileLogger:    nil,
	}
}

func (l *JLogger) EnableFileLog(enabled bool) {
	l.enableFileLog = enabled
}

func (l *JLogger) SetRotateType(rotateType RotateType) {
	l.rotateType = rotateType
}

func (l *JLogger) SetLogFile(filename string) {
	var (
		rtime time.Duration
		myFmt string
	)
	switch l.rotateType {
	case RotateMinute:
		myFmt = ".%Y-%m-%d-%H-%M"
		rtime = time.Minute
	case RotateHour:
		myFmt = ".%Y-%m-%d-%H"
		rtime = time.Hour
	case RotateDay:
		myFmt = ".%Y-%m-%d"
		rtime = time.Hour * 24
	default:
		myFmt = ".%Y-%m-%d"
		rtime = time.Hour * 24
	}

	hook, err := rotatelogs.New(
		filename+myFmt,
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(rtime),
	)
	if err != nil {
		panic(err)
	}

	fp, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	if l.fileLogger == nil {
		fileLogger := log.New(fp, "", log.LstdFlags)
		l.fileLogger = fileLogger
	}
	l.fileLogger.SetOutput(hook)

}

func (l *JLogger) Infof(format string, v ...interface{}) {
	mfmt := fmt.Sprintf("%c[1;40;32m%v%c[0m%v", 0x1B, InfoStr, 0x1B, format)
	l.stdLogger.Printf(mfmt, v...)
	if l.enableFileLog {
		l.fileLogger.Printf(format, v...)
	}
}

func (l *JLogger) Errorf(format string, v ...interface{}) {
	mfmt := fmt.Sprintf("%c[1;40;41m%v%c[0m%v", 0x1B, ErrorStr, 0x1B, format)
	l.stdLogger.Printf(mfmt, v...)
	if l.enableFileLog {
		l.fileLogger.Printf(format, v...)
	}
}

func (l *JLogger) InfoKV(kv KV, msg string) {
	kv["msg"] = msg
	bytes, _ := json.Marshal(kv)
	mfmt := fmt.Sprintf("%c[1;40;32m%v%c[0m%v", 0x1B, InfoStr, 0x1B, string(bytes))
	l.stdLogger.Printf(mfmt)
	if l.enableFileLog {
		l.fileLogger.Printf(string(bytes))
	}
}

func (l *JLogger) ErrorKV(kv KV, msg string) {
	kv["msg"] = msg
	bytes, _ := json.Marshal(kv)
	mfmt := fmt.Sprintf("%c[1;40;41m%v%c[0m%v", 0x1B, ErrorStr, 0x1B, string(bytes))
	l.stdLogger.Printf(mfmt)
	if l.enableFileLog {
		l.fileLogger.Printf(string(bytes))
	}
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func InfoKV(kv KV, msg string) {
	logger.InfoKV(kv, msg)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

func ErrorKV(kv KV, msg string) {
	logger.ErrorKV(kv, msg)
}

func EnableFileLog(enabled bool) {
	logger.EnableFileLog(enabled)
}

func SetFileLog(filename string) {
	logger.SetLogFile(filename)
}

func SetRotateType(rotateType RotateType) {
	logger.SetRotateType(rotateType)
}
