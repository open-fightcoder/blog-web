package controllers

import "github.com/gin-gonic/gin"

func Register(router *gin.RouterGroup) {
	RegisterSelf(router)
	RegisterUser(router)
	RegisterArticle(router)
}
