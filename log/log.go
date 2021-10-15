package log

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"

	"os"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type BuiltinLogger struct {
	logger *log.Entry
}

func NewBuiltinLogger() *BuiltinLogger {
	log.SetHandler(text.New(os.Stderr))
	log.SetLevel(log.InfoLevel)
	ctx := log.WithFields(log.Fields{
		"game": "Brick Break",
	})
	return &BuiltinLogger{logger: ctx}
}

func (l *BuiltinLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *BuiltinLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *BuiltinLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}
