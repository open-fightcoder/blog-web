package common

import (
	"github.com/shiyicode/gin-framework/common/g"
	"github.com/shiyicode/gin-framework/common/store"
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
