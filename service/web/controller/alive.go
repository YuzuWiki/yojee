package controller

import (
	"github.com/gin-gonic/gin"
)

type AliveController struct{}

func (ctrl *AliveController) Alive(ctx *gin.Context) {
	ctx.JSON(200, success())
}
