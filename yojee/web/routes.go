package web

import (
	"github.com/like9th/yojee/yojee/web/controller"
)

func (svr *Server) RegisterRoutes()  {
	// 健康检查 api
	alive := controller.AliveController{}
	svr.GET("/alive", alive.Alive)
}