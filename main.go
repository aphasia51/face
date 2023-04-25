package main

import "main/initialize"

func main() {
	// 初始化日志
	initialize.InitLogger()

	// 初始化配置
	initialize.InitConfig()

	// 初始化MYSQL连接
	initialize.ConnectMySQL()

}
