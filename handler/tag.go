package handler

import (
	"YouComic-Nano/datasource"
	"YouComic-Nano/serializer"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var TagListHandler gin.HandlerFunc = func(context *gin.Context) {
	page, pageSize := GetPagination(context)
	tags, count, err := datasource.GetTagList(page, pageSize)
	if err != nil {
		ServerError(err, context, 500)
		return
	}
	template := serializer.BaseTagTemplate{}
	logrus.Debug(context.Request.URL.String())
	data := serializer.SerializeMultipleTemplate(tags, &template, nil)
	container := serializer.DefaultListContainer{}
	container.SerializeList(data, map[string]interface{}{
		"page":     page,
		"pageSize": pageSize,
		"count":    count,
		"url":      context.Request.URL,
	})
	context.JSON(200, container)

}
