package models

import (
	"time"
)

// 获取上传时间
func TimeFile() (string, string) {
	// 获取时间戳
	t := time.Now().Format("2006-01-02 15:04:05")
	year := t[0:4]
	month := t[5:7]
	day := t[8:10]
	hour := t[11:13]
	minute := t[14:16]
	second := t[17:19]
	date := year + month + day
	file := date + hour + minute + second
	return date, file
}

func Judge(extName string) bool {
	allowExtMap := map[string]bool{".jpg": true, ".png": true, ".gif": true, ".jpeg": true, ".webp": true}
	if _, ok := allowExtMap[extName]; !ok {
		return false
	}
	return true
}
