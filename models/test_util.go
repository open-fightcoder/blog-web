package models

import (
	"sync"

	"github.com/shiyicode/gin-framework/common"
)

var once sync.Once

func InitAllInTest() {
	once.Do(func() {
		common.Init("../cfg/cfg.toml.debug")
	})
}
