package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// RawLogger a new instance of the logger. You can have any number of instances.
var RawLogger = logrus.New()

// Logger logger
var Logger = RawLogger.WithFields(logrus.Fields{"field1": "test1"})

func init() {
	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	RawLogger.Out = os.Stdout

	if IsProd() {
		RawLogger.Formatter = &logrus.JSONFormatter{}
		RawLogger.Level = logrus.InfoLevel
	} else {
		RawLogger.Level = logrus.DebugLevel
	}

	Logger.Infoln("logging is on...")
}
