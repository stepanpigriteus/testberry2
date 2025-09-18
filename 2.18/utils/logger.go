package utils

import (
	"log/slog"
	"os"

	"grep/2.18/domain"
)

type Slogger struct {
	logger *slog.Logger
}

func NewSlogger() domain.Logger {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	slogger := slog.New(handler)
	return &Slogger{logger: slogger}
}

func (l *Slogger) Info(msg string, args ...interface{}) {
	l.logger.Info(msg, args)
}

func (l *Slogger) Error(msg string, args ...interface{}) {
	l.logger.Error(msg, args)
}

func (l *Slogger) Warn(msg string, args ...interface{}) {
	l.logger.Warn(msg, args)
}

func (l *Slogger) Debug(msg string, args ...interface{}) {
	l.logger.Debug(msg, args)
}
