package controllers

import (
	"net/http"

	error "github.com/blog-web/common/g"
	"github.com/blog-web/managers"
	"github.com/blog-web/router/controllers/base"
	"github.com/gin-gonic/gin"
)

func RegisterArticle(router *gin.RouterGroup) {
	router.GET("article/love", httpHandlerLove)
	router.GET("article/unlove", httpHandlerUnlove)
	router.GET("article/listloves", httpHandlerListloves)
}

func httpHandlerLove(c *gin.Context) {
	articleId := c.Query("article_id")
	userId := base.UserId(c)
	if userId == 0 {
		panic(error.PrivError("您尚未登录!"))
	}
	managers.ArticleLove(userId, articleId)
	c.JSON(http.StatusOK, "")
}

func httpHandlerUnlove(c *gin.Context) {
	articleId := c.Query("article_id")
	userId := base.UserId(c)
	if userId == 0 {
		panic(error.PrivError("您尚未登录!"))
	}
	managers.ArticleUnlove(userId, articleId)
	c.JSON(http.StatusOK, "")
}

func httpHandlerListloves(c *gin.Context) {
	//TODO
}
