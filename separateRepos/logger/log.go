package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

var Namespace string

type defaultFieldFormatter func(*logrus.Entry) ([]byte, error)

// Format implements logrus.Formatter interface
func (f defaultFieldFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Data["name"] = Namespace
	return f(e)
}

func Init(namespace string) {
	Namespace = namespace
	formatter := logrus.JSONFormatter{TimestampFormat: time.RFC3339}

	logrus.SetFormatter(defaultFieldFormatter(
		formatter.Format,
	))
}
