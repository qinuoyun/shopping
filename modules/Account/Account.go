package Account

import (
	"fmt"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"gorm.io/gorm"
	"reflect"
)

type Account struct {
	gorm.Model
	Title    string
	CateName int
	State    int
}

// 初始化生成路由
func init() {
	path := Account{}
	ci.RegisterModule(&path, reflect.TypeOf(path).PkgPath())
}

// TableName 表示配置操作数据库的表名称
func (Account) TableName() string {
	return fmt.Sprintf("ci_%s", "account")
}
