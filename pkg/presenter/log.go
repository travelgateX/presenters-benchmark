package presenter

import (
	"github.com/travelgateX/go-io/log"
)

type Logger interface {
	Info(string)
}

type logger struct {
	*log.Logger
}

var _ Logger = (*logger)(nil)

func (l *logger) Info(msg string) {
	l.Logger.Info(msg)
}

func NewLogger(log *log.Logger) Logger {
	return &logger{
		Logger: log,
	}
}

func NewStdoutLogger() Logger {
	return &logger{
		Logger: log.NewStdLogger(),
	}
}
