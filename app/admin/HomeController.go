package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"github.com/qinuoyun/shopping/modules/Article"
	"reflect"
)

// Home 用于自动注册路由
type Home struct{}

// 初始化生成路由
func init() {
	path := Home{}
	ci.Register(&path, reflect.TypeOf(path).PkgPath())
}

// Index 方法
func (con Home) Index(c *gin.Context) {

	var articleList []Article.Article
	//
	//ci.DB.Preload("Article").Find(&articleList)
	//ci.GetDemo()

	ci.M("Article").Find(&articleList)

	c.JSON(200, gin.H{
		"result": "测阿基德",
		"data":   articleList,
	})
}
