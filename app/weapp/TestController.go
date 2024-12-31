package weapp

import (
	"github.com/gin-gonic/gin"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"reflect"
)

// 用于自动注册路由
type Test struct{}

// 初始化生成路由
func init() {
	fpath := Test{}
	ci.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// Index 方法
func (hc Test) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the home page of admin!",
	})
}
