package main

import (
	"fmt"
	"panLite/api"
	"panLite/dao"
	"panLite/module/database"
)

func main() {
	// 初始化数据库
	database.NewDatabase()
	row, err := dao.GetByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(row)
	// 初始化缓存
	// 初始化用户鉴权
	// 初始化路由
	r := api.NewRoute()
	r.Run(":8080")
}
