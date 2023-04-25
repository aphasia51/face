package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"face/api"
	"face/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")

	zap.S().Info("用户相关URL")

	{
		UserRouter.GET("login", api.Login)
		UserRouter.GET("register", api.Register)
	}
}

func InitCutRouter(Router *gin.RouterGroup) {
	CutRouter := Router.Group("func")

	CutRouter.Use(middleware.AuthJWTMiddleware())

	zap.S().Info("功能相关URL")

	{
		CutRouter.GET("commit", api.OneCut)
		CutRouter.GET("push", api.Notice)
		CutRouter.GET("download", api.DownLoad)
	}
}
