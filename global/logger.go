package global

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger = nil

func InitLogger() *zerolog.Logger {
	if Logger == nil {
		zerolog.TimeFieldFormat = "2006-01-02 15:04:05.000"

		output := zerolog.ConsoleWriter{Out: os.Stdout}
		output.FormatTimestamp = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[ %s ]", i))
		}
		output.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s:", i)
		}
		output.FormatFieldValue = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("%s", i))
		}

		logger := zerolog.New(output).With().Timestamp().Logger()
		Logger = &logger
	}

	return Logger
}

func CloseLogger() error {
	return nil
}
