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
	// Configure the logger
	zerolog.MessageFieldName = "log-message"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.LevelFieldName = "log-level"
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return fmt.Sprintf("%s:%d", file[strings.LastIndex(file, "/")+1:], line)
	}
}

// level is a wrapper on top of zerolog's level
type level zerolog.Level

// raw returns the underlying zerolog's level
func (l level) raw() zerolog.Level {
	return zerolog.Level(l)
}

const (
	// Selective log levels
	LevelDebug = level(zerolog.DebugLevel)
	LevelInfo  = level(zerolog.InfoLevel)
	LevelWarn  = level(zerolog.WarnLevel)

	// Default log level
	defaultLevel = LevelInfo
)

// Logger has selective logger method signatures
type Logger interface {
	Debug() *zerolog.Event
	Info() *zerolog.Event
	Warn() *zerolog.Event
	Error() *zerolog.Event
	Fatal() *zerolog.Event
	With() zerolog.Context
}

// key holds the type that is used as key to fetch value from context.
// The value in ctx will be a logger so the same logger instance
// can be retrived anywhere the context is available.
type key struct{}

// NewLogger creates a new logger with specified output stream and level.
// This returns the updated context which has key struct as key and logger as value.
// The same logger instance can be retrived from this context.
func NewLogger(ctx context.Context, out io.Writer, level level) (context.Context, Logger) {
	logger := newLogger(out, level)
	ctx = context.WithValue(ctx, key{}, logger)

	return ctx, logger
}

// FromContext retrives the logger from context.
// In absence of logger in context, this creates a new logger with Stdout
// as output stream and default log level specified in constants.
func FromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(key{}).(*zerolog.Logger); ok {
		return logger
	}

	return newLogger(os.Stdout, defaultLevel)
}

// newLogger internally creates a new logger instance enabling
// timestamp, caller info and specified level
func newLogger(out io.Writer, level level) Logger {
	logger := zerolog.New(out).With().
		Timestamp().Caller().Logger().
		Level(level.raw())

	return &logger
}

// SilentLogger creates a new logger that discards all the output.
// This doesn't serve the logger functionality,  but is used in places
// where the silent logger would replace the actual logger object,
// and suppress all the output.
func SilentLogger(ctx context.Context) (context.Context, Logger) {
	silentLogger := newLogger(io.Discard, defaultLevel)
	ctx = context.WithValue(ctx, key{}, silentLogger)

	return ctx, silentLogger
}

// FromRawLogger creates Logger type from zerolog Logger
func FromRawLogger(l zerolog.Logger) Logger {
	return &l
}
