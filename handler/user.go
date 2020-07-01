package handler

import "github.com/gin-gonic/gin"

var UserAuthHandler gin.HandlerFunc = func(context *gin.Context) {
	context.JSON(200, map[string]interface{}{
		"id":   1,
		"sign": "",
	})
}
