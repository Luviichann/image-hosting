package controllers

import (
	"fmt"
	"goimg/models"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

// 用于判断权限
func InitMiddleware(ctx *gin.Context) {
	//获取cookie
	_, err := ctx.Cookie("userId")
	// 判断有没有登录，有没有cookie
	if err != nil {
		// 如果没有cookie说明没有登录，那就先去登录。
		fmt.Println(err)
		ctx.Redirect(http.StatusFound, "/user/login")
	}
	fmt.Println("我是一个中间件")
}

// 上传页面
func (con AuthController) Upload(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.html", gin.H{})
}

// 上传逻辑
func (con AuthController) DoAdd(ctx *gin.Context) {
	//获取cookie
	userId, _ := ctx.Cookie("userId")

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
	num, _ := strconv.Atoi(userId)
	DBAdd(date, file.Filename, num)
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

func DBAdd(date, fileName string, userId int) {
	// 获取属性
	image := models.Image{
		Filename:   fileName,
		FileBelong: date,
		UserId:     userId,
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
	// fmt.Printf("count: %v\n", count)
	if count == 0 {
		if err := tx.Create(&folder).Error; err != nil {
			fmt.Printf("err: %v\n", err)
			tx.Rollback()
			return
		}
	}

	tx.Commit()
}

// 注册页面
func (con AuthController) Register(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", gin.H{})
}

// 注册逻辑
func (con AuthController) DoAddUser(ctx *gin.Context) {
	user := models.User{}
	// 从前台获取信息
	key := ctx.PostForm("key")
	// 检验key是不是真的
	var count int64
	models.DB.Model(&models.License{}).Where("secret_key=?", key).Count(&count)
	if count == 0 {
		// 假的
		return
	} else {
		// 真的，并且删除。
		license := models.License{}
		models.DB.Where("secret_key = ?", key).Delete(&license) //查询到secret_key=key的数据并删除。
	}
	// 注册
	if err := ctx.ShouldBind(&user); err == nil {
		user.Auth = 1
		fmt.Printf("user: %v\n", user)
		ctx.JSON(http.StatusOK, gin.H{
			"id":       user.Id,
			"username": user.Username, //""里的是返回值，:后的是变量
			"password": user.Password,
			"auth":     user.Auth,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
	if err := models.DB.Create(&user).Error; err != nil {
		fmt.Println("数据添加失败！")
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("数据添加成功！")
	}
}

// 登录页面
func (con AuthController) Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

// 登录逻辑
func (con AuthController) DoLogin(ctx *gin.Context) {
	// 获取登录的用户名和密码
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	// 判断用户名是否真的存在
	user := models.User{}
	models.DB.Where("username=?", username).Find(&user)
	fmt.Println(user)
	if user.Id == 0 {
		// 用户名不存在
		fmt.Println("用户名不存在")
		return
	} else {
		// 继续判断密码是否正确
		if user.Password != password {
			// 密码错误
			fmt.Println("密码错误")
			return
		}
	}
	// 登录成功
	//设置cookie，localhost是域名或者ip。
	ctx.SetCookie("userId", strconv.Itoa(user.Id), 3600, "/", "localhost", false, false)
}
