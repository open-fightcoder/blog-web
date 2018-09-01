package controllers

import (
	"net/http"

	error "github.com/blog-web/common/g"
	"github.com/blog-web/managers"
	"github.com/blog-web/models"
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
	comment := models.Comment{}
	err := c.Bind(&comment)
	if err != nil {
		panic(error.ParamError("参数格式错误,解析失败!"))
	}
	userId := base.UserId(c)
	if userId == 0 {
		panic(error.PrivError("您尚未登录!"))
	}
	managers.CommentAdd(userId, comment.ArticleId, comment.Content)
	c.JSON(http.StatusOK, "")
}

func httpHandlerCommentReply(c *gin.Context) {
	comment := models.Comment{}
	err := c.Bind(&comment)
	if err != nil {
		panic(error.ParamError("参数格式错误,解析失败!"))
	}
	userId := base.UserId(c)
	if userId == 0 {
		panic(error.PrivError("您尚未登录!"))
	}
	managers.CommentReply(userId, comment.Id, comment.ArticleId, comment.Content)
	c.JSON(http.StatusOK, "")
}

func httpHandlerCommentDelete(c *gin.Context) {
	commentId := c.Query("comment_id")
	userId := base.UserId(c)
	if userId == 0 {
		panic(error.PrivError("您尚未登录!"))
	}
	managers.CommentDelete(userId, commentId)
	c.JSON(http.StatusOK, "")
}

func httpHandlerCommentList(c *gin.Context) {
	commentId := c.Query("comment_id")
	userId := base.UserId(c)
	if userId == 0 {
		panic(error.PrivError("您尚未登录!"))
	}
	commentData := managers.CommentList(commentId)
	c.JSON(http.StatusOK, commentData)
}
