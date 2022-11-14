package main

import (
	"ginblog/model"
	"ginblog/routers"
)

func main() {
	// 引用数据库
	model.InitDb()
	// 引入redis
	model.InitRedis()

	routers.InitRouter()
}
