package controller

// import (
// 	"fmt"
// 	"strconv"
//
// 	"github.com/gin-gonic/gin"
//
// 	"github.com/YuzuWiki/yojee/global"
// 	"github.com/YuzuWiki/yojee/service/pixiv_service"
// )
//
// type PixivController struct{}
//
// func (ctr *PixivController) Sync(ctx *gin.Context) {
// 	phpsessid := ctx.Query("phpsessid")
//
// 	pid, err := strconv.ParseInt(ctx.Query("pid"), 10, 0)
// 	if len(phpsessid) == 0 || err != nil {
// 		ctx.JSON(400, fail(400, "Miss pid"))
// 		return
// 	}
// 	psrv := pixiv_service.NewService(phpsessid, 6, 10)
//
// 	go func() {
// 		if err := psrv.SyncUser(pid); err != nil {
// 			global.Logger.Err(err).Msg(fmt.Sprintf("[SyncUser] phpsessid=%s  pid=%d", phpsessid, pid))
// 		}
// 	}()
//
// 	go func() {
// 		if err := psrv.SyncArtworks(pid); err != nil {
// 			global.Logger.Err(err).Msg(fmt.Sprintf("[SyncArtworks] phpsessid=%s  pid=%d", phpsessid, pid))
// 		}
// 	}()
//
// 	ctx.JSON(200, success())
// }

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/YuzuWiki/yojee/service/pixiv_service"
)

type PixivController struct {
	srv pixiv_service.Service
}

func (ctr *PixivController) GetPid(ctx *gin.Context) {
	phpsessid := ctx.Query("phpsessid")

	pid, err := ctr.srv.GetPid(phpsessid)
	if err != nil {
		ctx.JSON(400, fail(400, err.Error()))
		return
	}

	ctx.JSON(200, success(map[string]int64{"pid": pid}))
	return
}

func (ctr *PixivController) Account(ctx *gin.Context) {
	pid, err := strconv.ParseInt(ctx.Query("pid"), 10, 64)
	if err != nil {
		ctx.JSON(400, fail(400, err.Error()))
		return
	}

	if data, err := ctr.srv.GetAccountInfo(pid); err == nil {
		ctx.JSON(200, success(data))
		return
	}

	if data, err := ctr.srv.FlushAccountInfo(pid); err != nil {
		ctx.JSON(400, fail(402, err.Error()))
		return
	} else {
		ctx.JSON(200, success(data))
	}
	return
}
