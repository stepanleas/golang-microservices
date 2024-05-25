package logger

import (
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger(hooks ...logrus.Hook) (*LogrusLogger, error) {
	log := logrus.New()

	return &LogrusLogger{logger: log}, nil
}

func (l *LogrusLogger) AddHook(hook logrus.Hook) {
	l.logger.AddHook(hook)
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LogrusLogger) Err(msg string, err error) {
	l.logger.Error(msg, logrus.WithField("error", err.Error()))
}
