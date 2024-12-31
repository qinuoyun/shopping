package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"reflect"
)

// 用于自动注册路由
type Home struct{}

// 初始化生成路由
func init() {
	fpath := Home{}
	ci.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// Index 方法
func (hc Home) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the home page of admin!",
	})
}
