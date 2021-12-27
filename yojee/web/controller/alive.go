package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/like9th/yojee/yojee/global"
)

type AliveController struct {
	baseController
}

func (ctrl *AliveController) Alive(ctx *gin.Context) {
	global.Logger.Info().Msg("这是一个测试信息")

	ctx.JSON(200, gin.H{
		"error":   0,
		"errmsg":  "success",
		"records": map[string]interface{}{},
	})
}
