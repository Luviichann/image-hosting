package routers

import (
	"goimg/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersInit(r *gin.Engine) {
	Routers := r.Group("/")
	{
		Routers.GET("/", controllers.Controller{}.Index)

		Routers.GET("/randomimage", controllers.Controller{}.RandomImage)
	}
}
