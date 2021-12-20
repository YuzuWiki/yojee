package global

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

var Logger *zerolog.Logger = nil

func InitLogger() *zerolog.Logger {
	once.Do(func() {
		if Logger != nil {
			return
		}
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat:  time.RFC3339}
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
	})
	return Logger
}

func  CloseLogger() error {
	return nil
}