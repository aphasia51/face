package main

import (
	"face/dao"
	"face/initialize"

	"go.uber.org/zap"
)

func main() {
	// 初始化日志
	initialize.InitLogger()

	// 初始化MYSQL配置
	dao.InitConfig()
	// 初始化MYSQL连接
	dao.ConnectMySQL()

	// 初始化路由
	Router := initialize.Routers()

	// 启动监听端口
	if err := Router.Run(); err != nil {
		zap.S().Panic("启动失败: ", err.Error())
	}
}
