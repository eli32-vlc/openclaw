package infra

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LogLevel represents a log level.
type LogLevel string

const (
	LogLevelSilent LogLevel = "silent"
	LogLevelFatal  LogLevel = "fatal"
	LogLevelError  LogLevel = "error"
	LogLevelWarn   LogLevel = "warn"
	LogLevelInfo   LogLevel = "info"
	LogLevelDebug  LogLevel = "debug"
	LogLevelTrace  LogLevel = "trace"
)

// InitLogging initializes the zerolog logger.
func InitLogging(level LogLevel) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	switch level {
	case LogLevelSilent:
		zerolog.SetGlobalLevel(zerolog.Disabled)
	case LogLevelFatal:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case LogLevelError:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case LogLevelWarn:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case LogLevelInfo:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case LogLevelDebug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case LogLevelTrace:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

// SubsystemLogger returns a logger with a subsystem field.
func SubsystemLogger(subsystem string) zerolog.Logger {
	return log.With().Str("subsystem", subsystem).Logger()
}
