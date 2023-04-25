package api

import (
	"fmt"
	"main/forms"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	loginForm := forms.LoginForm{
		UserName: "",
		PassWord: "",
	}
	if err := ctx.ShouldBind(&loginForm); err != nil {
		return
	}

}

func Register(ctx *gin.Context) {
	// 用户注册
	registerForm := forms.RegisterForm{
		UserName: "",
		PassWord: "",
	}
	if err := ctx.ShouldBind(&registerForm); err != nil {
		return
	}

	fmt.Println(registerForm)
}
