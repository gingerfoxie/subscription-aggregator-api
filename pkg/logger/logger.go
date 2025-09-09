package logger

import (
	"github.com/sirupsen/logrus"
)

func InitLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}
