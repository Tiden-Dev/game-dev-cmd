package cmd

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	logger *zerolog.Logger
)

func initLogging() *zerolog.Logger {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	zeroLogger := zerolog.New(consoleWriter).With().Timestamp().Logger()
	return &zeroLogger
}
