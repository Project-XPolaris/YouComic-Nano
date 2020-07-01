package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetIntQueryParam(key string, defaultValue int, ctx *gin.Context) (int, error) {
	rawValue := ctx.Query(key)
	if len(rawValue) == 0 {
		return defaultValue, nil
	}
	value, err := strconv.Atoi(rawValue)
	return value, err
}
