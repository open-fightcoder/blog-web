package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/TV4/graceful"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shiyicode/gin-framework/common"
	"github.com/shiyicode/gin-framework/common/g"
	"github.com/shiyicode/gin-framework/router"
)

func main() {
	cfgFile := flag.String("c", "cfg/cfg.toml.debug", "set config file")
	flag.Parse()

	common.Init(*cfgFile)
	defer common.Close()

	router := router.GetRouter()

	graceful.LogListenAndServe(&http.Server{
		Addr:    fmt.Sprintf(":%d", g.Conf().Run.HTTPPort),
		Handler: router,
	})

	common.Close()
}
