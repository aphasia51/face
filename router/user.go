package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"main/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")

	zap.S().Info("用户相关URL")

	{
		UserRouter.GET("login", api.Login)
		UserRouter.GET("register", api.Register)
	}
}
