package Account

import (
	"fmt"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"gorm.io/gorm"
	"reflect"
)

type Wechat struct {
	gorm.Model
	Title string
	Name  int
	State int
}

// 初始化生成路由
func init() {
	path := Wechat{}
	ci.RegisterModule(&path, reflect.TypeOf(path).PkgPath())
}

// TableName 表示配置操作数据库的表名称
func (Wechat) TableName() string {
	return fmt.Sprintf("ci_%s", "wechat")
}
