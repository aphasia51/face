package initialize

import (
	"net/http"

	"face/middleware"
	"face/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(middleware.Cors())

	Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Code":   http.StatusOK,
			"Status": "Success",
		})
	})

	CutGroup := Router.Group("/v1/cut")
	router.InitUserRouter(CutGroup)
	router.InitCutRouter(CutGroup)

	return Router
}
