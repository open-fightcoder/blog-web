package controllers

import (
	"net/http"

	error "github.com/blog-web/common/g"
	"github.com/blog-web/managers"
	"github.com/blog-web/router/controllers/base"
	"github.com/gin-gonic/gin"
)

func RegisterComment(router *gin.RouterGroup) {
	router.GET("comment/add", httpHandlerCommentAdd)
	router.GET("comment/reply", httpHandlerCommentReply)
	router.GET("comment/delete", httpHandlerCommentDelete)
	router.GET("comment/list", httpHandlerCommentList)
}

func httpHandlerCommentAdd(c *gin.Context) {
	articleId := c.Query("article_id")
	userId := base.UserId(c)
	if userId == 0 {
		panic(error.PrivError("您尚未登录!"))
	}
	managers.ArticleLove(userId, articleId)
	c.JSON(http.StatusOK, "")
}

func httpHandlerCommentReply(c *gin.Context) {
	articleId := c.Query("article_id")
	userId := base.UserId(c)
	if userId == 0 {
		panic(error.PrivError("您尚未登录!"))
	}
	managers.ArticleUnlove(userId, articleId)
	c.JSON(http.StatusOK, "")
}

func httpHandlerCommentDelete(c *gin.Context) {

}

func httpHandlerCommentList(c *gin.Context) {

}
