package handler

import (
	"YouComic-Nano/config"
	"YouComic-Nano/datasource"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
	"strings"
)

var BookContentHandler gin.HandlerFunc = func(context *gin.Context) {
	var err error
	rawBookId := context.Param("id")
	bookId, err := strconv.Atoi(rawBookId)
	if err != nil {
		ServerError(err, context, 500)
		return
	}

	book, err := datasource.GetBookById(bookId)
	if err != nil {
		ServerError(err, context, 500)
		return
	}
	filename := context.Param("filename")

	if strings.Contains(filename, "cover_thumbnail") {

		thumbnail, err := datasource.GetBookCoverCacheById(book.ID)
		if err != nil {
			ServerError(err, context, 500)
			return
		}
		if len(thumbnail) != 0 {
			context.File(filepath.Join(config.RootPath, ".cache", "thumbnail", thumbnail))
			return
		} else {
			filename = book.Cover
			context.File(filepath.Join(config.RootPath, book.Path, filename))
			return
		}

	}
	context.File(filepath.Join(config.RootPath, book.Path, filename))
}
