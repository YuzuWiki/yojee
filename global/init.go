package global

func init() {
	InitLogger()

	Logger.Debug().Msg("global init.....")
	initRedis()
	InitDB()
	InitCron()
	initPixiv()

	Logger.Debug().Msg("global init finish")
}
