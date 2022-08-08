package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
)

type PixivController struct {
	baseController
}

func (ctr *PixivController) Sync(ctx *gin.Context) {
	phpsessid := ctx.Query("phpsessid")
	if len(phpsessid) == 0 {
		ctx.JSON(400, fail(400, "Miss ssesid"))
		return
	}

	go func(s string) {
		pCtx := pixiv.NewContext(s)

		pid, err := pCtx.Pid()
		if err != nil {
			global.Logger.Error().Msg(fmt.Sprintf("UserInfo Error: invalid phpsessid, %s", s))
			return
		}

		db := global.DB()

		var user model.PixivUserMod
		if !db.Raw(
			"SELECT user.id AS id FROM pixiv_user AS user WHERE user.pid = ? LIMIT 1;", pid,
		).Scan(&user).RecordNotFound() {
			data, err := apis.InfoAPI{}.Info(pCtx, pid)
			if err != nil {
				global.Logger.Error().Msg(fmt.Sprintf(" UserInfo Error: (%s) error=%s", phpsessid, err.Error()))
				return
			}

		}

		// 获取账户信息

		tx := db.Begin()

	}(phpsessid)

	ctx.JSON(200, success())
	return
}
