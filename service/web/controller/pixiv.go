package controller

import (
	"fmt"
	"os"
	"strconv"

	"github.com/YuzuWiki/Pixivlee/apis"
	"github.com/gin-gonic/gin"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv"
)

type PixivController struct {
	srv pixiv.Service
}

func (ctr *PixivController) GetPid(ctx *gin.Context) {
	phpsessid := os.Getenv("PIXIV_PHPSESSID")
	if len(phpsessid) == 0 {
		ctx.JSON(400, fail(400, "invalid phpsessid"))
		return
	}

	pid, err := ctr.srv.GetPid(phpsessid)
	if err != nil {
		ctx.JSON(400, fail(400, err.Error()))
		return
	}

	ctx.JSON(200, success(map[string]int64{"pid": pid}))
	return
}

func (ctr *PixivController) Account(ctx *gin.Context) {
	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
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
	} else {
		ctx.JSON(200, success(data))
	}
	return
}

func (ctr *PixivController) FlushAccount(ctx *gin.Context) {
	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		ctx.JSON(400, fail(400, err.Error()))
		return
	}

	if data, err := ctr.srv.FlushAccountInfo(pid); err != nil {
		ctx.JSON(400, fail(402, err.Error()))
	} else {
		ctx.JSON(200, success(data))
	}
	return
}

func (ctr *PixivController) GetFollowing(ctx *gin.Context) {
	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		ctx.JSON(400, fail(400, "invalid pid"))
		return
	}

	params := struct {
		Pid    int64 `json:"pid"`
		Limit  int   `json:"limit"`
		Offset int   `json:"offset" `
	}{
		Pid:    pid,
		Limit:  24,
		Offset: 0,
	}
	if err = ctx.ShouldBind(&params); err != nil {
		ctx.JSON(400, fail(400, fmt.Sprintf("params error, %s", err.Error())))
		return
	}

	follows, err := ctr.srv.GetFollowing(params.Pid, params.Limit, params.Offset)
	if err != nil {
		ctx.JSON(400, fail(400, fmt.Sprintf("get following error, %s", err.Error())))
		return
	}

	ctx.JSON(200, success(follows))
	return
}

func (ctr *PixivController) SyncFollowing(ctx *gin.Context) {
	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		ctx.JSON(400, fail(400, err.Error()))
		return
	}

	go func() {
		_, _ = ctr.srv.SyncFollowing(pid)
	}()
	ctx.JSON(200, success())
}

func (ctr *PixivController) GetArtWork(ctx *gin.Context) {
	params := struct {
		ArtType string `form:"art_type"`
		ArtId   int64  `form:"art_id"`
	}{}
	if err := ctx.BindQuery(&params); err != nil {
		ctx.JSON(400, fail(400, err.Error()))
		return
	}

	artWork, err := ctr.srv.GetArtwork(params.ArtType, params.ArtId)
	if err != nil {
		ctx.JSON(400, fail(400, err.Error()))
	} else {
		ctx.JSON(200, success(artWork))
	}
	return
}

func (ctr *PixivController) GetArtWorks(ctx *gin.Context) {
	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		ctx.JSON(400, fail(400, err.Error()))
		return
	}

	params := struct {
		Pid     int64
		ArtType string `form:"art_type"`
		Limit   int    `form:"limit"`
		Offset  int    `form:"offset"`
	}{
		Pid:     pid,
		ArtType: apis.Illust,
		Limit:   24,
		Offset:  0,
	}
	if err = ctx.BindQuery(&params); err != nil {
		ctx.JSON(400, fail(400, err.Error()))
		return
	}

	artWorks, err := ctr.srv.GetArtworks(params.Pid, params.ArtType, params.Limit, params.Offset)
	if err != nil {
		ctx.JSON(400, fail(400, err.Error()))
	} else {
		ctx.JSON(200, success(artWorks))
	}
	return
}

func (ctr *PixivController) SyncArtWorks(ctx *gin.Context) {
	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		ctx.JSON(400, fail(400, err.Error()))
		return
	}

	go func() {
		if err := ctr.srv.SyncArtWorks(pid); err != nil {
			global.Logger.Error().Msg(fmt.Sprintf("[SyncArtWorks] (%9d): ERROR, errmsg=%s", pid, err.Error()))
		} else {
			global.Logger.Info().Msg(fmt.Sprintf("[SyncArtWorks] (%9d): SUCCESS", pid))
		}
	}()

	ctx.JSON(200, success())
	return
}
