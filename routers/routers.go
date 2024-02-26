package routers

import (
	"goimg/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersInit(r *gin.Engine) {
	Routers := r.Group("/")
	{
		Routers.GET("/", controllers.Controller{}.Index)
		Routers.GET("/upload", controllers.Controller{}.Upload)
		Routers.POST("/doAdd", controllers.Controller{}.DoAdd)
		Routers.GET("/randomimage", controllers.Controller{}.RandomImage)
	}
}
