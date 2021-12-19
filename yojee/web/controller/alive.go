package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/like9th/yojee/yojee/interfaces"
)

type AliveController struct {
	baseController
}

func (ctrl *AliveController) Alive(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"error":  interfaces.Logger.GetLevel(),
		"errmsg": fmt.Sprintf("logger level",),
		"records": map[string]interface{}{},
	})
}
