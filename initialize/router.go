package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Code":   http.StatusOK,
			"Status": "Success",
		})
	})

	return Router
}
