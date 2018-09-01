package controllers

import (
	"net/http"

	error "github.com/blog-web/common/g"
	"github.com/blog-web/managers"
	"github.com/blog-web/router/controllers/base"
	"github.com/gin-gonic/gin"
)

func RegisterArticle(router *gin.RouterGroup) {
	router.POST("articleType/writer",httpHandlerArticleTypeWriter)
	router.GET("articleType/reader",httpHandlerArticleTypeReader)
	router.POST("article/writer",httpHandlerArticleWriter)
	router.GET("article/reader",httpHandlerArticleReader)
	router.GET("article/love", httpHandlerLove)
	router.GET("article/unlove", httpHandlerUnlove)
	router.GET("article/listloves", httpHandlerListloves)
}
func httpHandlerArticleTypeWriter(c *gin.Context){
	userId := base.UserId(c)
	if userId == 0 {
		panic(error.PrivError("您尚未登录!"))
	}
	indexStr := c.PostForm("index")
	articleTypeInfo:=c.PostForm("articleTypeInfo")
	managers.ArticleTypeWriter(userId,indexStr,articleTypeInfo)
	c.JSON(http.StatusOK, "")
}
func httpHandlerArticleTypeReader(c *gin.Context){
	articleTypeIdStr:=c.Query("articleTypeId")
	articleType:=managers.ArticleTypeReader(articleTypeIdStr)
	c.JSON(http.StatusOK, gin.H{
		"index":   articleType.Index,
		"articleTypeInfo": articleType.ArticleTypeInfo,
	})
}
func httpHandlerArticleReader(c *gin.Context){
	articleTypeIdStr:=c.Query("articleId")
	article := managers.ArticleReader(articleTypeIdStr)
	c.JSON(http.StatusOK, gin.H{
		"title":   article.Title,
		"content": article.Content,
		"status":    article.Status,
		"articleTypeId":    article.ArticleType,
		"look":    article.Look,
		"favour":    article.Favour,
	})
}
func httpHandlerArticleWriter(c *gin.Context){
	userId := base.UserId(c)
	if userId == 0 {
		panic(error.PrivError("您尚未登录!"))
	}
	title :=c.PostForm("title")
	content:=c.PostForm("content")
	statusStr:=c.PostForm("status")
	articleTypeIdStr:=c.PostForm("article_type")
	lookStr:=c.PostForm("look")
	favourStr:=c.PostForm("favour")
	managers.ArticleWriter(userId,title,content,statusStr,articleTypeIdStr,lookStr,	favourStr)
	c.JSON(http.StatusOK, "")
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
