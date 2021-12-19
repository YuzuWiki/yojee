package main

import (
	"github.com/like9th/yojee/yojee/web"
)

func main()  {
	svr := web.NewServer(8081)
	_ = svr.Run()
}
