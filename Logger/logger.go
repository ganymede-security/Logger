package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

func New(debugMode bool) *Logger {
	toDebug := zerolog.InfoLevel
	if debugMode {
		toDebug = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel((toDebug))
	newLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	return &Logger{logger: &newLogger}
}
