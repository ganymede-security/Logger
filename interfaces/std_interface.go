package interfaces

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Set global log value
func SetGlobalLevel(level string) {
	input := strings.ToLower(level)
	switch input {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

}

// stores messages to log later, from standard interface
type Event struct {
	id      int
	message string
}

// enforces specific log message formats
type StandardLogger struct {
	logger *zerolog.Logger
}

// Initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = zerolog.New(os.Stderr).With().Timestamp().Logger()

	var standardLogger = &StandardLogger{logger: &baseLogger}

	return standardLogger
}

// declare variables to store log messages as new Events
var (
	invalidArgMessage      = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage = Event{2, "Invalid value for argument: "}
	missingArgMessage      = Event{3, "Missing arg: %s"}
)

// standard error message
func (l *StandardLogger) InvalidArg(argumentName string) {
	l.logger.Error().Msg((invalidArgMessage.message + " " + argumentName))
}

// standard error message
func (l *StandardLogger) InvalidArgValue(argumentName string, T any) {
	l.logger.Error().Msg(fmt.Sprint(invalidArgValueMessage.message+" "+argumentName+" ", T))
}

// standard error message
func (l *StandardLogger) MissingArg(argumentName string) {
	l.logger.Error().Msg((missingArgMessage.message + " " + argumentName))
}

func (l *StandardLogger) Std(T any) {
	l.logger.Info().Stack().Msg(fmt.Sprint(T))
}
