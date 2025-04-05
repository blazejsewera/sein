package monitor

import (
	"log/slog"
	"os"
)

type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
	Fatal(string, ...any)
}

type logger struct {
	*slog.Logger
}

func (l *logger) Fatal(msg string, args ...any) {
	l.Error(msg, args...)
	panic(msg)
}

var log *logger

func init() {
	w := os.Stderr
	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	log = &logger{slog.New(slog.NewTextHandler(w, opts))}
}

func Log() Logger {
	return log
}
