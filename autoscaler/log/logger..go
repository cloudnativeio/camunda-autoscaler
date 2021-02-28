package log

import (
"os"
"time"

"github.com/sirupsen/logrus"
)

type GopherLogger struct {
	*logrus.Logger
}

func NewLogger() *GopherLogger {
	var baseLogger = logrus.New()

	var gopherLogger = &GopherLogger{baseLogger}
	gopherLogger.Level = logrus.DebugLevel
	gopherLogger.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
		},
		TimestampFormat: time.RFC3339,
	}
	gopherLogger.Out = os.Stdout

	return gopherLogger
}

