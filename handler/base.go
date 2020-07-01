package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	DEFAULT_PAGE_SIZE = 20
	DEFAULT_PAGE      = 1
)

func GetPagination(ctx *gin.Context) (int, int) {
	var err error
	page := DEFAULT_PAGE
	pageSize := DEFAULT_PAGE_SIZE
	pageStr := ctx.Query("page")
	if len(pageStr) > 0 {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return page, pageSize
		}
	}
	pageSizeStr := ctx.Query("page_size")
	if len(pageSizeStr) > 0 {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			return page, pageSize
		}
	}

	return page, pageSize
}
