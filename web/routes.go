package web

import (
	"github.com/YuzuWiki/yojee/web/controller"
)

func (svr *Server) RegisterRoutes() {
	// 健康检查 api
	alive := controller.AliveController{}
	svr.GET("/alive", alive.Alive)

	// pixiv := controller.PixivController{}
	// svr.GET("/pixiv/user/Sync", pixiv.Sync)

	pixiv := controller.PixivController{}
	svr.GET("/pixiv/account/pid", pixiv.GetPid)
	svr.GET("/pixiv/account/info", pixiv.Account)
	svr.POST("/pixiv/account/following", pixiv.GetFollowing)
	svr.PUT("/pixiv/account/following", pixiv.SyncFollowing)
	svr.GET("/pixiv/artwork/info", pixiv.GetArtWork)
}
