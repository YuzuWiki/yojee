package controller

import (
	"github.com/gin-gonic/gin"
)

type AliveController struct {
	baseController
}

func (ctrl *AliveController) Alive(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"error":  0,
		"errmsg": "success",
	})
}
