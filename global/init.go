package global

func init() {
	InitLogger()

	Logger.Debug().Msg("global init.....")
	initRedis()
	InitDB()
	InitCron()
	initPool()

	Logger.Debug().Msg("global init finish")
}
