package controllers

import (
	"goimg/models"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (con Controller) Index(ctx *gin.Context) {
	// ctx.String(http.StatusOK, "首页")
	images := []models.Image{}
	models.DB.Order("id Desc").Find(&images)
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"images": images,
	})
}

func (con Controller) RandomImage(ctx *gin.Context) {
	images := []models.Image{}
	models.DB.Order("id Desc").Find(&images)
	image := images[rand.Intn(len(images))]
	src := "/img/" + image.FileBelong + "/" + image.Filename
	ctx.Redirect(http.StatusFound, src)
}
