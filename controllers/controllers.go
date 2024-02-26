package controllers

import (
	"fmt"
	"goimg/models"
	"math/rand"
	"net/http"
	"os"
	"path"

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

func (con Controller) Upload(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.html", gin.H{})
}

func (con Controller) DoAdd(ctx *gin.Context) {
	//获取文件
	file, err := ctx.FormFile("image")
	extName := path.Ext(file.Filename)
	//判断文件名是否合法
	if !models.Judge(extName) {
		ctx.String(200, "文件类型不合法")
		return
	}
	//返回上传信息
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// 上传文件到指定的目录

	//获取日期文件夹
	var date string
	//获取文件名
	date, file.Filename = models.TimeFile()
	//加上后缀
	file.Filename += extName

	//写入数据库
	DBAdd(date, file.Filename)
	//创建图片保存目录
	dir := "./img/" + date
	if err := os.MkdirAll(dir, 0666); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	dst := path.Join(dir, file.Filename)
	fmt.Println(dst)
	ctx.SaveUploadedFile(file, dst)
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("'%s' uploaded!", file.Filename)})
}

func (con Controller) RandomImage(ctx *gin.Context) {
	images := []models.Image{}
	models.DB.Order("id Desc").Find(&images)
	image := images[rand.Intn(len(images))]
	src := "/img/" + image.FileBelong + "/" + image.Filename
	ctx.Redirect(http.StatusFound, src)
}

func DBAdd(date, fileName string) {
	// 获取属性
	image := models.Image{
		Filename:   fileName,
		FileBelong: date,
	}
	folder := models.Folder{
		FolderName: date,
	}
	// 开启事务
	tx := models.DB.Begin()
	if err := tx.Create(&image).Error; err != nil {
		fmt.Printf("err: %v\n", err)
		tx.Rollback()
		return
	}
	// 判断日期文件夹是否已经存在，如果存在就不再向数据库里存。
	var count int64
	models.DB.Model(&models.Folder{}).Where("folder_name=?", date).Count(&count)
	fmt.Printf("count: %v\n", count)
	if count == 0 {
		if err := tx.Create(&folder).Error; err != nil {
			fmt.Printf("err: %v\n", err)
			tx.Rollback()
			return
		}
	}

	tx.Commit()
}
