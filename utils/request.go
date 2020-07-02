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

func GetIntArrayQueryParam(key string, ctx *gin.Context) ([]int, error) {
	rawValues := ctx.QueryArray(key)
	if len(rawValues) == 0 {
		return make([]int, 0), nil
	}
	result := make([]int, 0)
	for _, rawValue := range rawValues {
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}

	return result, nil
}
