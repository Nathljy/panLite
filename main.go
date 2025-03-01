package main

import "panLite/api"

func main() {
	// 初始化数据库
	// 初始化缓存
	// 初始化用户鉴权
	// 初始化路由
	r := api.NewRoute()
	r.Run(":8080")
}
