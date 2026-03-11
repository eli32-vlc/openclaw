package logging

import (
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	mu          sync.Mutex
	logFile     *os.File
	initialized bool
)

// InitLogging initializes the logger with the given level and optional file path.
func InitLogging(level string, logFilePath string, console bool) {
	mu.Lock()
	defer mu.Unlock()

	writers := []io.Writer{}

	if console {
		writers = append(writers, zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		})
	}

	if logFilePath != "" {
		if err := os.MkdirAll(filepath.Dir(logFilePath), 0o700); err == nil {
			f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
			if err == nil {
				logFile = f
				writers = append(writers, f)
			}
		}
	}

	var writer io.Writer
	if len(writers) == 0 {
		writer = io.Discard
	} else if len(writers) == 1 {
		writer = writers[0]
	} else {
		writer = zerolog.MultiLevelWriter(writers...)
	}

	log.Logger = zerolog.New(writer).With().Timestamp().Logger()
	setLevel(level)
	initialized = true
}

func setLevel(level string) {
	switch level {
	case "silent", "off":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn", "warning":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

// SubsystemLogger returns a logger tagged with the given subsystem.
func SubsystemLogger(subsystem string) zerolog.Logger {
	return log.With().Str("subsystem", subsystem).Logger()
}

// Close closes any open log files.
func Close() {
	mu.Lock()
	defer mu.Unlock()
	if logFile != nil {
		logFile.Close()
		logFile = nil
	}
}
