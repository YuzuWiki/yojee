package global

func init() {
	InitLogger()

	Logger.Debug().Msg("global init.....")
	InitDB()
	InitCron()

	Logger.Debug().Msg("global init finish")
}
