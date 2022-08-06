package controller

import (
	"github.com/gin-gonic/gin"
)

type baseController struct {
}

// Success return body
func success(data ...any) gin.H {
	if len(data) == 0 {
		return gin.H{
			"errcode": 0,
			"errmsg":  "success",
		}
	} else {
		return gin.H{
			"errcode": 0,
			"errmsg":  "success",
			"data":    data[0],
		}
	}

}

// fail return body
func fail(code int, errMsg string) gin.H {
	return gin.H{
		"errcode": code,
		"err_msg": errMsg,
	}
}
