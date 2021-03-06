package controllers

import (
	"net/http"

	"github.com/blog-web/common/g"
	"github.com/gin-gonic/gin"
)

func RegisterSelf(router *gin.RouterGroup) {
	router.GET("self/health", httpHandlerHealth)
	router.GET("self/config", httpHandlerConfig)
	router.GET("self/reload", httpHandlerReload)
}

func httpHandlerHealth(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func httpHandlerConfig(c *gin.Context) {
	c.JSON(http.StatusOK, g.Conf())
}

func httpHandlerReload(c *gin.Context) {
	g.LoadConfig(g.ConfigFile)
	c.String(http.StatusOK, "reload succeed")
}
