package global

func init() {
	InitLogger()

	Logger.Debug().Msg("global init.....")
	initRedis()
	InitDB()
	InitCron()
	initPixiv()
	initPool()

	Logger.Debug().Msg("global init finish")
}
