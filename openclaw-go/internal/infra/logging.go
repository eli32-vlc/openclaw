package infra

import (
	"github.com/openclaw/openclaw-go/internal/logging"
	"github.com/rs/zerolog"
)

// LogLevel represents a log level string understood by InitLogging.
type LogLevel = string

const (
	LogLevelSilent LogLevel = "silent"
	LogLevelFatal  LogLevel = "fatal"
	LogLevelError  LogLevel = "error"
	LogLevelWarn   LogLevel = "warn"
	LogLevelInfo   LogLevel = "info"
	LogLevelDebug  LogLevel = "debug"
	LogLevelTrace  LogLevel = "trace"
)

// InitLogging initializes the logger. It delegates to the logging package
// so there is a single source of truth for log initialization.
func InitLogging(level LogLevel) {
	logging.InitLogging(level, "", true)
}

// SubsystemLogger returns a logger tagged with the given subsystem.
// It delegates to the logging package.
func SubsystemLogger(subsystem string) zerolog.Logger {
	return logging.SubsystemLogger(subsystem)
}
