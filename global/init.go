package global

func init() {
	InitLogger()

	Logger.Debug().Msg("global init.....")
	initRedis()
	InitDB()
	InitCron()

	Logger.Debug().Msg("global init finish")
}
