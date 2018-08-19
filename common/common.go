package common

import (
	"github.com/blog-web/common/g"
	"github.com/blog-web/common/store"
)

func Init(cfgFile string) {
	g.LoadConfig(cfgFile)
	g.InitLog()
	store.InitMysql()
	store.InitRedis()
}

func Close() {
	store.CloseMysql()
	store.CloseRedis()
}
