package main

import (
	"fmt"
	"gotask/phaseOne/task3/code"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gormproject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&code.Student{}, &code.Transaction{},
		&code.Account{}, &code.Employee{},
		&code.User{}, &code.Post{}, &code.Comment{})
	// fmt.Printf("添加了%d条学生信息\n", code.CreateStuInfo(db))
	// var students []code.Student = code.ConditionalQueryStuInfo(db)
	// fmt.Println(students)
	// var student code.Student = code.UpdateStuInfo(db, "Charlie", 6)
	// fmt.Println(student)
	// fmt.Printf("删除了%d条学生信息\n", code.ConditionalDeleteStuInfo(db, "age", "<", 15))
	// fmt.Printf("添加了%d条账户信息\n", code.CreateAccountInfo(db))
	// fmt.Println(code.AccounTrading(db, 1, 2, 100))

	dbsqlx, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gormproject?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	//插入数据还是用gorm框架实现   查询使用sqlx 来实现
	// fmt.Printf("添加了%d条职工信息\n", code.CreateEmployInfo(db))
	//技术部的字典项 是 1
	fmt.Println(code.QueryDepartmentEmplInfo(dbsqlx, 1))
	fmt.Println(code.QueryMaxSalaryEmployee(dbsqlx))
	//
	var price float64 = 50
	fmt.Printf("价格大于%.2f价格的书籍\n", price)
	fmt.Println(code.QueryExpensiveBooks(dbsqlx, price))

	fmt.Println(code.QueryPostDetailInfo(db, 2))
	fmt.Println(code.QueryMaxCommentForPost(db))
	//新增一个文章 然后看钩子函数会不会触发
	// newPost := &code.Post{
	// 	Title:   "新文章标题",
	// 	Content: "这里是文章内容",
	// 	UserID:  1,
	// }
	// code.PublishPost(db, newPost)
	// comment1 := &code.Comment{Content: "写得很好！", PostID: 1, UserID: 1}
	// comment2 := &code.Comment{Content: "学习了，感谢分享。", PostID: 1, UserID: 2}
	// comment3 := &code.Comment{Content: "有点疑问，能详细说说吗？", PostID: 2, UserID: 3}
	// comment4 := &code.Comment{Content: "支持作者！", PostID: 2, UserID: 1}
	// comment5 := &code.Comment{Content: "内容很实用。", PostID: 3, UserID: 2}
	// code.MakeComment(db, comment1)
	// code.MakeComment(db, comment2)
	// code.MakeComment(db, comment3)
	// code.MakeComment(db, comment4)
	// code.MakeComment(db, comment5)
	code.DelectComment(db, 27)
}
