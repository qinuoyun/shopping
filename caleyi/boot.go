package caleyi

import (
	"fmt"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"strings"
)

func BootStart() {
	//加载路由
	r := InitRouter()

	fmt.Printf("执行了BootStart需要下不")

	routes := ""
	for _, route := range r.Routes() {

		fmt.Printf("日志	route.Path：%v\n", route.Path)
		fmt.Printf("日志A：%v\n", !strings.Contains(route.Path, "/*filepath"))

		if !strings.Contains(route.Path, "/admin/") && route.Path != "/" && !strings.Contains(route.Path, "/*filepath") {
			routes = routes + fmt.Sprintf("%v\n", route.Path)
		}
	}
	filePath := "runtime/app/routers.txt"
	ci.WriteToFile(filePath, routes)

	err := r.Run(":9097")
	if err != nil {
		return
	}
}
