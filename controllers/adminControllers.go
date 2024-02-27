package controllers

import (
	"goimg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

// 判断权限
func JudgeAdmin(ctx *gin.Context) {
	// fmt.Printf("userId: %v\n", userId)
	//获取cookie
	userId, _ := ctx.Get("userId")
	user := models.User{}
	models.DB.Where("id = ?", userId).Find(&user)
	if user.Auth <= 1 {
		// 权限不够
		// 重定向到首页
		ctx.Redirect(http.StatusFound, "/")
	}
}

func (con AdminController) Admin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin.html", gin.H{})
}
