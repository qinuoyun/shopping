package Article

import (
	"fmt"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"gorm.io/gorm"
	"reflect"
)

type Article struct {
	gorm.Model
	Title   string
	Content int
	State   int
}

// 初始化生成路由
func init() {
	path := Article{}
	ci.RegisterModule(&path, reflect.TypeOf(path).PkgPath())
}

// TableName 表示配置操作数据库的表名称
func (Article) TableName() string {
	return fmt.Sprintf("ci_%s", "article")
}
