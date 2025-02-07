package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"github.com/qinuoyun/shopping/modules/Account"
	"github.com/qinuoyun/shopping/modules/Article"
	"net/http"
	"reflect"
)

// Login 用于自动注册路由
type Login struct{}

// 初始化生成路由
func init() {
	path := Login{}
	ci.Register(&path, reflect.TypeOf(path).PkgPath())
}

// Index 方法
func (con Login) Index(c *gin.Context) {

	var account Account.Account
	// 将请求体中的 JSON 数据绑定到 account 结构体上
	if err := c.ShouldBindJSON(&account); err != nil {
		// 如果绑定出错，返回 400 错误及错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("打印", account)

	// 在此处可以将账号信息存储到数据库或其他存储介质中

	// 打印接收到的账号信息
	c.JSON(http.StatusOK, gin.H{"data": account})
}

func (con Login) Register(c *gin.Context) {

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
