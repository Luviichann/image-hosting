package routers

import (
	"goimg/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutersInit(r *gin.Engine) {
	Routers := r.Group("/user")
	{
		Routers.GET("/upload", controllers.IsLogin, controllers.AuthController{}.Upload)
		Routers.POST("/doAddImage", controllers.AuthController{}.DoAdd)

		Routers.GET("/register", controllers.AuthController{}.Register)
		Routers.POST("/doAddUser", controllers.AuthController{}.DoAddUser)

		Routers.GET("/login", controllers.AuthController{}.Login)
		Routers.POST("/doLogin", controllers.AuthController{}.DoLogin)

		Routers.GET("/admin", controllers.IsLogin, controllers.JudgeAdmin, controllers.AdminController{}.Admin)
	}
}
