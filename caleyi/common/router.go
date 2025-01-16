package common

//一定要导入这个Controller包，用来注册需要访问的方法
//这里路由-由构架是添加-开发者仅在指定工程目录下controller.go文件添加宝即可
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"time"
)

func InitRouter() *gin.Engine {
	//初始化路由
	R := gin.Default()
	R.SetTrustedProxies([]string{"127.0.0.1"})
	//访问公共目录
	R.Static("/public", "./public")
	//访问域名根目录重定向
	R.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "message": "欢迎使用卡莱易框架"})
	})

	R.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization", "Businessid", "verify-encrypt", "ignoreCancelToken", "verify-time"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//5.找不到路由
	R.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		c.JSON(404, gin.H{"code": 404, "message": "您" + method + "请求地址：" + path + "不存在！"})
	})
	//绑定基本路由，访问路径：/User/List
	ci.Bind(R)
	return R
}
