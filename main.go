package main

import (
	_ "github.com/qinuoyun/shopping/app"
	"github.com/qinuoyun/shopping/caleyi"
	_ "github.com/qinuoyun/shopping/modules"
)

func main() {

	// 启动服务器
	caleyi.BootStart()
}
