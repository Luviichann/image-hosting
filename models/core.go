package models

import (
	"fmt"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	cfg, _ := ini.Load("./config.ini")
	IP := cfg.Section("mysql").Key("ip").String()
	PORT := cfg.Section("mysql").Key("port").String()
	USER := cfg.Section("mysql").Key("user").String()
	PASSWORD := cfg.Section("mysql").Key("password").String()
	DATABASE := cfg.Section("mysql").Key("database").String()
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "root:root@tcp(127.0.0.1:3306)/goimage?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", USER, PASSWORD, IP, PORT, DATABASE)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Println("数据库连接失败！")
	} else {
		fmt.Println("数据库连接成功！")
	}
}
