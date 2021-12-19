package interfaces

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger = nil

func InitLogger() *zerolog.Logger {
	once.Do(func() {
		if Logger != nil {
			return
		}

		//*Logger = zerolog.New(os.Stderr)
		Logger = &zerolog.Logger{}
		Logger.Output(os.Stdin)
	})
	return Logger
}

func  CloseLogger() error {
	return nil
}