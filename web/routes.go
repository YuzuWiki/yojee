package web

import (
	"yojee/web/controller"
)

func (svr *Server) RegisterRoutes() {
	// 健康检查 api
	alive := controller.AliveController{}
	svr.GET("/alive", alive.Alive)
}
