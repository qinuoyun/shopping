package Account

import (
	"fmt"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"reflect"
)

// AccountStatus 账号状态枚举
type AccountStatus int

const (
	// Active 表示账号处于激活状态
	Active AccountStatus = iota
	// Inactive 表示账号处于未激活状态
	Inactive
	// Suspended 表示账号处于暂停状态
	Suspended
)

// Account 账号模型
type Account struct {
	// Username 存储用户名，json 标签用于序列化和反序列化，binding 标签用于验证，要求必填且长度在 3 到 20 之间
	Username string `json:"username" binding:"required,min=3,max=20"`
	// Password 存储用户密码，json 标签用于序列化和反序列化，binding 标签用于验证，要求必填且长度至少为 6
	Password string `json:"password" binding:"required,min=6"`
	// Phone 存储用户手机号，json 标签用于序列化和反序列化，binding 标签用于验证，要求必填且需通过自定义 phone 验证器验证
	Phone string `json:"phone" binding:"required,phone"`
	// Status 存储账号状态，json 标签用于序列化和反序列化，binding 标签用于验证，要求必填
	Status AccountStatus `json:"status" binding:"required"`
	// FullName 存储用户的全名，json 标签用于序列化和反序列化，binding 标签用于验证，要求必填
	FullName string `json:"full_name" binding:"required"`
	// Gender 存储用户性别，json 标签用于序列化和反序列化，binding 标签用于验证，要求必填且只能是 male、female 或 other 中的一个
	Gender string `json:"gender" binding:"required,oneof=male female other"`
	// Email 存储用户邮箱，json 标签用于序列化和反序列化，binding 标签用于验证，要求必填且需为有效的邮箱地址
	Email string `json:"email" binding:"required,email"`
	// CreatedAt 存储账号创建时间，json 标签用于序列化和反序列化
	CreatedAt string `json:"created_at"`
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
