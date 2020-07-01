package handler

import "github.com/gin-gonic/gin"

func ServerError(err error, ctx *gin.Context, code int) {
	ctx.SecureJSON(code, err.Error())
}
