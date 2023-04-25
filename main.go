package main

import (
	"main/initialize"

	"go.uber.org/zap"
)

func main() {
	// 初始化日志
	initialize.InitLogger()

	// 初始化配置
	initialize.InitConfig()

	// 初始化MYSQL连接
	initialize.ConnectMySQL()

	// 初始化路由
	Router := initialize.Routers()

	if err := Router.Run(); err != nil {
		zap.S().Panic("启动失败: ", err.Error())
	}
}
