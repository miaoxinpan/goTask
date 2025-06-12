package main

import (
	"gotask/phaseOne/task4/blogPorject/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB() // 初始化数据库
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
