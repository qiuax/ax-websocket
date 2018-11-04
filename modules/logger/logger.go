package logger

import (
	"ax-websocket/helper"
	"ax-websocket/modules/mongodb"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	logger.Formatter = &logrus.JSONFormatter{}

	// 添加mongodbHooker
	logger.AddHook(&mongodbHooker{Session: mongodb.SessionLogs})

	// 调试级别
	//logger.Level = logrus.DebugLevel
	//logger.Out = os.Stdout

	logger.Out = ioutil.Discard
}

// Fields wraps logrus.Fields, which is a map[string]interface{}
type Fields logrus.Fields

func IsDebugLevel() bool {
	if logger.Level == logrus.DebugLevel {
		return true
	}

	return false
}

func GetLogger() *logrus.Logger {
	return logger
}

func SetLogLevel(level logrus.Level) {
	logger.Level = level
}

func SetLogFormatter(formatter logrus.Formatter) {
	logger.Formatter = formatter
}

// 添加额外字段
func addExtendField(entry *logrus.Entry) {
	entry.Data["file"] = fileInfo(3)
	entry.Data["ip"] = helper.LocalIp

	entry.Data["nanotime"] = helper.GetTimeNano()
	//entry.Data["timestamp"] = helper.GetTimestamp()
	entry.Data["datetime"] = helper.ChinaTimeNow().Format(time.RFC3339)
}

// Debug logs a message with fields at level Debug on the standard logger.
func Debugf(f Fields, format string, args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields(f))
		addExtendField(entry)
		entry.Debugf(format, args...)
	}
}

// Debug logs a message with fields at level Debug on the standard logger.
func Infof(f Fields, format string, args ...interface{}) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields(f))
		addExtendField(entry)
		entry.Infof(format, args...)
	}
}

// Debug logs a message with fields at level Debug on the standard logger.
func Warnf(f Fields, format string, args ...interface{}) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields(f))
		addExtendField(entry)
		entry.Warnf(format, args...)
	}
}

// Debug logs a message with fields at level Debug on the standard logger.
func Errorf(f Fields, format string, args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields(f))
		addExtendField(entry)
		entry.Errorf(format, args...)
	}
}

// Debug logs a message with fields at level Debug on the standard logger.
func Fatalf(f Fields, format string, args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields(f))
		addExtendField(entry)
		entry.Fatalf(format, args...)
	}
}

// Debug logs a message with fields at level Debug on the standard logger.
func Panicf(f Fields, format string, args ...interface{}) {
	if logger.Level >= logrus.PanicLevel {
		entry := logger.WithFields(logrus.Fields(f))
		addExtendField(entry)
		entry.Panicf(format, args...)
	}
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		file = truncateFile(file)
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// 将/data/go/src/aliserver/modules/nats/nats.go截取为aliserver/modules/nats/nats.go
func truncateFile(file string) string {
	flag := "/src/"

	pos := strings.Index(file, flag)
	if pos > -1 {
		return file[pos+len(flag):]
	}

	return file
}
