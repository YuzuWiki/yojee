package global

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

// init, load config from .env (file name)
func init() {
	_ = godotenv.Load(".env")
}

// init global (outside) services
func init() {
	InitLogger()
	Logger.Debug().Msg("[init] Debug: global services initializing")

	initRedis()
	InitDB()

	Logger.Debug().Msg("[init] Debug: global services done")
}

// init global (internal) components
func init() {
	Logger.Debug().Msg("[init] Debug: global components initializing")
	InitCron()
	initPool()
	Logger.Debug().Msg("[init] Debug: global components done ")
}
