package router

import (
	"sync"

	"github.com/blog-web/common/g"
	"github.com/blog-web/router/controllers/apiv1"
	"github.com/blog-web/router/middleware"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var once sync.Once

// 获取路由并初始化
func GetRouter() *gin.Engine {
	// 只执行一次
	once.Do(func() {
		initRouter()
	})
	return router
}

// 初始化路由
func initRouter() {
	router = gin.New()
	gin.SetMode(g.Conf().Run.Mode)

	router.Use(middleware.Logger())
	router.Use(middleware.Cors())
	router.Use(middleware.Recovery())
	router.Use(middleware.MaxAllowed(g.Conf().Run.MaxAllowed))

	apiRouter := router.Group("/apiv1")
	apiv1.Register(apiRouter)
}
