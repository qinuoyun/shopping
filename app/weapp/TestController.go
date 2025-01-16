package weapp

import (
	"github.com/gin-gonic/gin"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"reflect"
)

// Test 用于自动注册路由
type Test struct{}

// 初始化生成路由
func init() {
	fPath := Test{}
	ci.Register(&fPath, reflect.TypeOf(fPath).PkgPath())
}

// Index 方法
func (hc Test) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": "Welcome to the home page of test!",
	})
}
