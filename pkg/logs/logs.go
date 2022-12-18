package logs

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.MessageFieldName = "log-message"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.LevelFieldName = "log-level"
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return fmt.Sprintf("%s:%d", file[strings.LastIndex(file, "/")+1:], line)
	}
}

type Logger interface {
	Info() *zerolog.Event
	Error() *zerolog.Event
	Fatal() *zerolog.Event
	With() zerolog.Context
}

type key struct{}

func FromContext(ctx context.Context) Logger {
	if l, ok := ctx.Value(key{}).(*zerolog.Logger); ok {
		return l
	}

	return newLogger(os.Stdout)
}

func SilentLogger(ctx context.Context) (context.Context, Logger) {
	silentLogger := newLogger(io.Discard)
	ctx = context.WithValue(ctx, key{}, silentLogger)

	return ctx, silentLogger
}

func newLogger(out io.Writer) Logger {
	logger := zerolog.New(out).With().Timestamp().Caller().Logger()
	return &logger
}
