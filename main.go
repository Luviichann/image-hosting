package main

import (
	"goimg/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/img", "./img")
	routers.RoutersInit(r)
	routers.AuthRoutersInit(r)
	r.Run()
}

// github.com/pilu/fresh
