package api

import (
	"face/dao"
	"face/forms"
	"face/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var db = dao.DB
	userForm := forms.UserForm{}
	// 检测表单数据是否完整
	if err := ctx.ShouldBind(&userForm); err != nil {
		return
	}

	// 查询用户是否存在
	check_user := db.Where("user_name = ?", userForm.UserName).
		First(&userForm)
	if check_user.RowsAffected >= 1 {
		ctx.Set("username", userForm.UserName)
		token, _ := middleware.GenerateToken(userForm.UserName)
		ctx.JSON(http.StatusOK, gin.H{
			"user": userForm.UserName,
			"msg":  "登录成功",
			"token": gin.H{
				"token": token,
			},
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "用户不存在, 请先注册",
		})
		return
	}
}

func Register(ctx *gin.Context) {
	var db = dao.DB
	// 用户注册
	userForm := forms.UserForm{}
	if err := ctx.ShouldBind(&userForm); err != nil {
		return
	}

	check_user := db.Where("user_name = ?", userForm.UserName).
		First(&userForm)
	if check_user.RowsAffected >= 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"user": userForm.UserName,
			"msg":  "用户存在，请直接登录",
		})
		return
	}

	user := forms.UserForm{
		UserName: userForm.UserName,
		PassWord: userForm.PassWord,
	}
	res := db.Create(&user)
	if res.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "创建用户失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "创建用户成功，请登录",
	})
}
