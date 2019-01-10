package presenter

import (
	"encoding/json"
	"github.com/travelgateX/go-io/log"
	"strings"
)

type Logger interface {
	Info(string)
	LogResult(Result)
}

type logger struct {
	*log.Logger
}

var _ Logger = (*logger)(nil)

func (l *logger) Info(msg string) {
	l.Logger.Info(msg)
}

func (l *logger) LogResult(r Result) {
	sb := strings.Builder{}
	json.NewEncoder(&sb).Encode(r)
	l.Logger.Info(sb.String())
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
