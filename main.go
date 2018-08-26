package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/TV4/graceful"
	"github.com/blog-web/common"
	"github.com/blog-web/common/g"
	"github.com/blog-web/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfgFile := flag.String("c", "/Users/leeezm/code/go/src/github.com/blog-web/cfg/cfg.toml.debug", "set config file")
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
