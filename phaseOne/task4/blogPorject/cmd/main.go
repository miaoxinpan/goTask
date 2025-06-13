package main

import (
	"fmt"
	"gotask/phaseOne/task4/blogPorject/config"
	"gotask/phaseOne/task4/blogPorject/router"
	"gotask/phaseOne/task4/blogPorject/structs"
)

func main() {
	config.InitDB() // 初始化数据库
	config.DB.AutoMigrate(&structs.Comment{}, &structs.Post{},
		&structs.User{})
	router := router.SetupRouter()
	//createData()
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务

}

func createData() {
	// 用户测试数据
	users := []structs.User{
		{Username: "alice", Password: "pass123", Email: "alice@example.com"},
		{Username: "bob", Password: "pass456", Email: "bob@example.com"},
		{Username: "charlie", Password: "pass789", Email: "charlie@example.com"},
	}

	// 文章测试数据
	posts := []structs.Post{
		{Title: "Go 入门", Content: "Go 是一门很棒的语言。", UserID: 1},
		{Title: "Gin 框架实践", Content: "Gin 用起来很方便。", UserID: 2},
		{Title: "GORM 使用技巧", Content: "GORM 支持自动迁移。", UserID: 1},
	}

	// 评论测试数据
	comments := []structs.Comment{
		{Content: "写得很好！", UserID: 2, PostID: 1},
		{Content: "有收获，感谢！", UserID: 3, PostID: 1},
		{Content: "学习了！", UserID: 1, PostID: 2},
		{Content: "有点疑问。", UserID: 2, PostID: 3},
	}
	for _, u := range users {
		config.DB.Create(&u)
		fmt.Println("插入没啊1？")
	}
	for _, p := range posts {
		config.DB.Create(&p)
		fmt.Println("插入没啊2？")
	}
	for _, c := range comments {
		config.DB.Create(&c)
		fmt.Println("插入没啊3？")
	}
}
