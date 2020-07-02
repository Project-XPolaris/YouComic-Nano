package handler

import (
	"YouComic-Nano/datasource"
	"YouComic-Nano/serializer"
	"YouComic-Nano/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

var BookListHandler gin.HandlerFunc = func(context *gin.Context) {

	page, pageSize := GetPagination(context)
	reader := datasource.BookReader{
		Page:     page,
		PageSize: pageSize,
	}
	// tag filter
	tagIds, err := utils.GetIntArrayQueryParam("tag", context)
	if err != nil {
		ServerError(err, context, 500)
		return
	}
	if len(tagIds) > 0 {
		reader.Filters = append(reader.Filters, &datasource.BookTagFilter{TagIds: tagIds})
	}

	nameSearch := context.Query("nameSearch")
	if len(nameSearch) > 0 {
		reader.Filters = append(reader.Filters, &datasource.BookNameSearchFilter{
			Key: nameSearch,
		})
	}
	logrus.Debug(context.Request.URL.String())
	books, count, err := reader.Read()
	if err != nil {
		ServerError(err, context, 500)
		return
	}

	template := serializer.BaseBookTemplate{}
	data := serializer.SerializeMultipleTemplate(books, &template, nil)
	container := serializer.DefaultListContainer{}
	container.SerializeList(data, map[string]interface{}{
		"page":     page,
		"pageSize": pageSize,
		"count":    count,
		"url":      context.Request.URL,
	})
	context.JSON(200, container)
}

var BookHandler gin.HandlerFunc = func(context *gin.Context) {
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

	template := serializer.BaseBookTemplate{}
	err = template.Serializer(*book, nil)
	if err != nil {
		ServerError(err, context, 500)
		return
	}

	context.JSON(200, template)
}

var BookTagHandler gin.HandlerFunc = func(context *gin.Context) {
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

	template := serializer.BaseTagTemplate{}
	data := serializer.SerializeMultipleTemplate(book.Tags, &template, nil)
	container := serializer.DefaultListContainer{}
	container.SerializeList(data, map[string]interface{}{
		"page":     1,
		"pageSize": 10,
		"count":    len(book.Tags),
		"url":      context.Request.URL,
	})
	context.JSON(200, container)
}
