package ci

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// ValidatePhone validatePhone 验证手机号的函数
func ValidatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	// 简单的手机号验证正则表达式，可根据实际需求修改
	re := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(phone)
}
