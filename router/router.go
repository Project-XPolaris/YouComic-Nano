package router

import (
	"YouComic-Nano/handler"
	"github.com/gin-gonic/gin"
)

func SetRouter(engine *gin.Engine) {
	engine.GET("books", handler.BookListHandler)
	engine.GET("pages", handler.PageListHandler)
	engine.POST("/user/auth", handler.UserAuthHandler)
	engine.GET("/content/book/:id/:filename", handler.BookContentHandler)
	engine.GET("/book/:id", handler.BookHandler)
	engine.GET("/book/:id/tags", handler.BookTagHandler)
	engine.GET("/tags", handler.TagListHandler)
}
