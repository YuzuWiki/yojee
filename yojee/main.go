package main

import (

	"github.com/like9th/yojee/yojee/interfaces"
	"github.com/like9th/yojee/yojee/web"
)

func main()  {
	interfaces.Logger.Error().Msg(">>>>>>>>>>>>>>>>>>>")
	svr := web.NewServer(8081)
	_ = svr.Run()
	interfaces.Logger.Error().Msg("<<<<<<<<<<<<<<<<<<<")
}
