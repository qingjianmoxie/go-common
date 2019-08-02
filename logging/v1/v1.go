package v1

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var logging *JLogger

type KV = logrus.Fields

func init() {
	logging = NewJLogger()
}

type JLogger struct {
	StdLogger     *logrus.Logger
	FileLogger    *logrus.Logger
	enableFileLog bool
}

func NewJLogger() *JLogger {
	stdLogger := logrus.New()
	stdLogger.SetFormatter(&logrus.TextFormatter{})
	stdLogger.SetLevel(logrus.DebugLevel)

	fileLogger := logrus.New()
	fileLogger.SetFormatter(&logrus.JSONFormatter{})
	fileLogger.SetLevel(logrus.InfoLevel)

	logger := &JLogger{
		StdLogger:     stdLogger,
		FileLogger:    fileLogger,
		enableFileLog: false,
	}
	return logger
}

func (l *JLogger) EnableFileLog(enabled bool) {
	l.enableFileLog = enabled
}

func (l *JLogger) SetLogfile(filename string) {
	l.FileLogger.Out = &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    50,
		MaxBackups: 7,
		MaxAge:     1,
		Compress:   true,
		LocalTime:  true,
	}
}

func (l *JLogger) Printf(format string, v ...interface{}) {
	l.StdLogger.Printf(format, v...)
	if l.enableFileLog {
		l.FileLogger.Printf(format, v...)
	}
}

func (l *JLogger) Debugf(format string, v ...interface{}) {
	l.StdLogger.Debugf(format, v...)
	if l.enableFileLog {
		l.FileLogger.Debugf(format, v...)
	}
}

func (l *JLogger) Infof(format string, v ...interface{}) {
	l.StdLogger.Infof(format, v...)
	if l.enableFileLog {
		l.FileLogger.Infof(format, v...)
	}
}

func (l *JLogger) Warnf(format string, v ...interface{}) {
	l.StdLogger.Warnf(format, v...)
	if l.enableFileLog {
		l.FileLogger.Warnf(format, v...)
	}
}

func (l *JLogger) Errorf(format string, v ...interface{}) {
	l.StdLogger.Errorf(format, v...)
	if l.enableFileLog {
		l.FileLogger.Errorf(format, v...)
	}
}

func (l *JLogger) Panicf(format string, v ...interface{}) {
	l.StdLogger.Panicf(format, v...)
	if l.enableFileLog {
		l.FileLogger.Panicf(format, v...)
	}
}

func (l *JLogger) Fatalf(format string, v ...interface{}) {
	l.StdLogger.Fatalf(format, v...)
	if l.enableFileLog {
		l.FileLogger.Fatalf(format, v...)
	}
}

func (l *JLogger) PrintKV(kv KV, msg string) {
	l.StdLogger.WithFields(kv).Printf(msg)
	if l.enableFileLog {
		l.FileLogger.WithFields(kv).Printf(msg)
	}
}

func (l *JLogger) DebugKV(kv KV, msg string) {
	l.StdLogger.WithFields(kv).Printf(msg)
	if l.enableFileLog {
		l.FileLogger.WithFields(kv).Printf(msg)
	}
}

func (l *JLogger) InfoKV(kv KV, msg string) {
	l.StdLogger.WithFields(kv).Infof(msg)
	if l.enableFileLog {
		l.FileLogger.WithFields(kv).Infof(msg)
	}
}

func (l *JLogger) WarnKV(kv KV, msg string) {
	l.StdLogger.WithFields(kv).Warnf(msg)
	if l.enableFileLog {
		l.FileLogger.WithFields(kv).Warnf(msg)
	}
}

func (l *JLogger) ErrorKV(kv KV, msg string) {
	l.StdLogger.WithFields(kv).Errorf(msg)
	if l.enableFileLog {
		l.FileLogger.WithFields(kv).Errorf(msg)
	}
}

func (l *JLogger) PanicKV(kv KV, msg string) {
	l.StdLogger.WithFields(kv).Panicf(msg)
	if l.enableFileLog {
		l.FileLogger.WithFields(kv).Panicf(msg)
	}
}

func (l *JLogger) FatalKV(kv KV, msg string) {
	l.StdLogger.WithFields(kv).Fatalf(msg)
	if l.enableFileLog {
		l.FileLogger.WithFields(kv).Fatalf(msg)
	}
}

func Printf(format string, v ...interface{}) {
	logging.Printf(format, v...)
}

func Debugf(format string, v ...interface{}) {
	logging.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	logging.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	logging.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logging.Errorf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	logging.Panicf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	logging.Fatalf(format, v...)
}

func PrintKV(kv KV, msg string) {
	logging.PrintKV(kv, msg)
}

func DebugKV(kv KV, msg string) {
	logging.DebugKV(kv, msg)
}

func InfoKV(kv KV, msg string) {
	logging.InfoKV(kv, msg)
}

func WarnKV(kv KV, msg string) {
	logging.WarnKV(kv, msg)
}

func ErrorKV(kv KV, msg string) {
	logging.ErrorKV(kv, msg)
}

func PanicKV(kv KV, msg string) {
	logging.PanicKV(kv, msg)
}

func FatalKV(kv KV, msg string) {
	logging.FatalKV(kv, msg)
}

func EnableLogFile(enabled bool) {
	logging.EnableFileLog(enabled)
}

func SetLogfile(filename string) {
	logging.SetLogfile(filename)
}
