package web

import (
	"github.com/gin-gonic/gin"

	"github.com/YuzuWiki/yojee/web/controller"
)

func (svr *Server) RegisterRoutes() {
	// 健康检查 api
	alive := controller.AliveController{}
	svr.router.GET("/alive", alive.Alive)

	// pixiv := controller.PixivController{}
	// svr.GET("/pixiv/user/Sync", pixiv.Sync)

	registerPixiv(svr.router.Group("/pixiv"))
}

func registerPixiv(group *gin.RouterGroup) {
	pixiv := controller.PixivController{}

	UserGroup := group.Group("/user")
	{
		UserGroup.GET("/phpsessid", pixiv.GetPid)
		UserGroup.GET("/:pid/info", pixiv.Account)
		UserGroup.POST("/:pid/following", pixiv.GetFollowing)
		UserGroup.PUT("/:pid/following", pixiv.SyncFollowing)
		UserGroup.GET("/:pid/artworks", pixiv.GetArtWorks)
		UserGroup.PUT("/:pid/artworks", pixiv.SyncArtWorks)
	}

	ArtworkGroup := group.Group("/artwork")
	{
		ArtworkGroup.GET("/:artId", pixiv.GetArtWork)
	}
}
