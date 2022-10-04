package web

import (
	controller2 "github.com/YuzuWiki/yojee/service/web/controller"
	"github.com/gin-gonic/gin"
)

func (svr *Server) RegisterRoutes() {
	// 健康检查 api
	alive := controller2.AliveController{}
	svr.router.GET("/alive", alive.Alive)

	registerPixiv(svr.router.Group("/pixiv"))
}

func registerPixiv(group *gin.RouterGroup) {
	pixiv := controller2.PixivController{}

	UserGroup := group.Group("/user")
	{
		UserGroup.GET("/phpsessid", pixiv.GetPid)
		UserGroup.GET("/:pid/info", pixiv.Account)
		UserGroup.PUT("/:pid/info", pixiv.FlushAccount)
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
