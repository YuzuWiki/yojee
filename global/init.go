package global

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	_ = godotenv.Load(".env")
}

func init() {
	// init some service

	InitLogger()

	Logger.Debug().Msg("global init.....")
	initRedis()
	InitDB()
	InitCron()
	initPool()

	Logger.Debug().Msg("global init finish")
}
