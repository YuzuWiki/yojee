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
