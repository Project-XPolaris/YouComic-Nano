package handler

import (
	"YouComic-Nano/datasource"
	"YouComic-Nano/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

var PageListHandler gin.HandlerFunc = func(context *gin.Context) {
	rawBookId := context.Query("book")
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
	templateParam := context.Query("template")
	var template serializer.TemplateSerializer
	template = &serializer.BasePageTemplate{}
	if templateParam == "withSize" {
		template = &serializer.PageTemplateWithSize{}
	}
	data := serializer.SerializeMultipleTemplate(book.Pages, template, map[string]interface{}{
		"bookId": book.ID,
	})

	container := serializer.DefaultListContainer{}
	container.SerializeList(data, map[string]interface{}{
		"page":     1,
		"pageSize": 999,
		"count":    len(book.Pages),
		"url":      context.Request.URL,
	})
	context.JSON(200, container)
}
