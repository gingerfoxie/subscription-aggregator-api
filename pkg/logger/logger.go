package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger(logOutput, logLevel string) {

	if logOutput == "" {
		logOutput = "stdout" // значение по умолчанию
	}

	// Настраиваем вывод в зависимости от значения
	if logOutput == "file" {
		// Создаем директорию для логов
		os.MkdirAll("logs", 0755)

		// Открываем файл для логов
		logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logrus.SetOutput(logFile)
		} else {
			logrus.Warn("Failed to log to file, using default stderr")
			// Оставляем stdout если файл не открылся
		}
	}
	// Если LOG_OUTPUT="stdout" или другое значение, оставляем стандартный вывод

	// Устанавливаем формат логов
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Устанавливаем уровень логирования из env (опционально)
	if logLevel == "" {
		logLevel = "info"
	}

	level, err := logrus.ParseLevel(logLevel)
	if err == nil {
		logrus.SetLevel(level)
	} else {
		logrus.SetLevel(logrus.InfoLevel) // уровень по умолчанию
	}
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
